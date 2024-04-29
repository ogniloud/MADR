package deck

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/ogniloud/madr/internal/db"
	"github.com/ogniloud/madr/internal/flashcards/models"
)

type Storage struct {
	Conn *db.PSQLDatabase
}

func (d *Storage) GetDecksByUserId(ctx context.Context, id models.UserId) (models.Decks, error) {
	rows, err := d.Conn.Query(ctx, `SELECT * FROM deck_config WHERE user_id=$1`, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	decks := models.Decks{}

	cfg := models.DeckConfig{}
	_, err = pgx.ForEachRow(rows, []any{&cfg.DeckId, &cfg.UserId, &cfg.Name}, func() error {
		decks[cfg.DeckId] = cfg

		return nil
	})

	if err != nil {
		return nil, err
	}

	return decks, nil
}

func (d *Storage) GetFlashcardsIdByDeckId(ctx context.Context, id models.DeckId) ([]models.FlashcardId, error) {
	rows, err := d.Conn.Query(ctx, `SELECT card_id FROM flashcard WHERE deck_id=$1`, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	ids := make([]models.FlashcardId, 0)

	cardId := 0
	_, err = pgx.ForEachRow(rows, []any{&cardId}, func() error {
		ids = append(ids, cardId)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return ids[:len(ids):len(ids)], nil
}

func (d *Storage) GetFlashcardById(ctx context.Context, id models.FlashcardId) (models.Flashcard, error) {
	row := d.Conn.QueryRow(ctx, `SELECT * FROM flashcard WHERE card_id=$1`, id)

	flashcard := models.Flashcard{}

	err := row.Scan(&flashcard.Id, &flashcard.W, &flashcard.B, &flashcard.DeckId, &flashcard.A)
	if err != nil {
		return models.Flashcard{}, fmt.Errorf("psql error: %w", err)
	}

	return flashcard, nil
}

func (d *Storage) GetLeitnerByUserIdCardId(ctx context.Context, id models.UserId, flashcardId models.FlashcardId) (models.UserLeitner, error) {
	row := d.Conn.QueryRow(ctx, `SELECT * FROM user_leitner WHERE user_id=$1 AND card_id=$2`, id, flashcardId)

	l := models.UserLeitner{}

	t := time.Time{}
	err := row.Scan(&l.Id, &l.UserId, &l.FlashcardId, &l.Box, &t)
	if err != nil {
		return models.UserLeitner{}, fmt.Errorf("psql error: %w", err)
	}

	l.CoolDown = models.CoolDown(t)

	return l, nil
}

func (d *Storage) GetUserInfo(ctx context.Context, uid models.UserId) (models.UserInfo, error) {
	row := d.Conn.QueryRow(ctx, `SELECT * FROM user_info WHERE user_id=$1`, uid)

	i := models.UserInfo{}

	err := row.Scan(&i.UserId, &i.MaxBox)
	if err != nil {
		return models.UserInfo{}, fmt.Errorf("psql error: %w", err)
	}

	return i, nil
}

func (d *Storage) PutAllFlashcards(ctx context.Context, id models.DeckId, cards []models.Flashcard) ([]models.FlashcardId, error) {
	batch := &pgx.Batch{}

	for _, v := range cards {
		batch.Queue(`INSERT INTO flashcard (word, backside, deck_id, answer) 
								VALUES ($1, $2, $3, $4) RETURNING card_id;`,
			v.W, v.B, id, v.A,
		)
	}

	results := d.Conn.SendBatch(ctx, batch)
	defer func() { _ = results.Close() }()

	rows, err := results.Query()
	if err != nil {
		return nil, fmt.Errorf("batch exec error: %w", err)
	}
	defer rows.Close()

	temp := 0
	ids := make([]models.FlashcardId, 0, len(cards))
	_, err = pgx.ForEachRow(rows, []any{&temp}, func() error {
		ids = append(ids, temp)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (d *Storage) PutNewDeck(ctx context.Context, config models.DeckConfig) (models.DeckId, error) {
	row := d.Conn.QueryRow(ctx,
		`INSERT INTO deck_config (user_id, name) VALUES ($1, $2) RETURNING deck_id`,
		config.UserId, config.Name,
	)

	id := 0

	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("psql error: %w", err)
	}

	return id, nil
}

func (d *Storage) PutAllUserLeitner(ctx context.Context, uls []models.UserLeitner) ([]models.LeitnerId, error) {
	batch := &pgx.Batch{}

	for _, v := range uls {
		batch.Queue(`INSERT INTO user_leitner (user_id, card_id, box, cool_down) 
								VALUES ($1, $2, $3, $4) RETURNING leitner_id;`,
			v.UserId, v.FlashcardId, v.Box, time.Time(v.CoolDown),
		)
	}

	results := d.Conn.SendBatch(ctx, batch)
	defer func() { _ = results.Close() }()

	rows, err := results.Query()
	if err != nil {
		return nil, fmt.Errorf("psql error: %w", err)
	}
	defer rows.Close()

	temp := 0
	ids := make([]models.LeitnerId, 0, len(uls))
	_, err = pgx.ForEachRow(rows, []any{&temp}, func() error {
		ids = append(ids, temp)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (d *Storage) DeleteFlashcardFromDeck(ctx context.Context, cardId models.FlashcardId) error {
	tx, err := d.Conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("delete card transaction failed: %w", err)
	}

	_, err = tx.Exec(ctx, `DELETE FROM user_leitner WHERE card_id=$1`, cardId)
	if err != nil {
		d.Conn.Logger().Errorf("delete leitners failed: %v, rollback...", err)
		defer func() {
			if err := tx.Rollback(ctx); err != nil {
				d.Conn.Logger().Errorf("rollback failed: %v", err)
			}
		}()

		return err
	}

	row := tx.QueryRow(ctx,
		`DELETE FROM flashcard WHERE card_id=$1 RETURNING card_id`, cardId)

	if err := row.Scan(&cardId); err != nil && !errors.Is(err, pgx.ErrNoRows) {
		d.Conn.Logger().Errorf("delete flashcards failed: %v, rollback...", err)
		defer func() {
			if err := tx.Rollback(ctx); err != nil {
				d.Conn.Logger().Errorf("rollback failed: %v", err)
			}
		}()

		return err
	}

	if err := tx.Commit(ctx); err != nil {
		d.Conn.Logger().Errorf("commit failed: %v", err)
		return fmt.Errorf("commit failed: %v", err)
	}

	return nil
}

func (d *Storage) DeleteUserDeck(ctx context.Context, userId models.UserId, deckId models.DeckId) error {
	tx, err := d.Conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("delete user deck tx begin failed: %w", err)
	}

	_, err = tx.Exec(ctx, `DELETE FROM user_leitner USING flashcard 
       WHERE flashcard.deck_id=$1 AND user_leitner.user_id=$2 AND flashcard.card_id=user_leitner.card_id`, deckId, userId)
	if err != nil {
		return fmt.Errorf("delete leitners failed: %w", err)
	}

	_, err = tx.Exec(ctx, `DELETE FROM flashcard WHERE flashcard.deck_id=$1`, deckId)
	if err != nil {
		return fmt.Errorf("delete falshcards failed: %w", err)
	}

	_, err = d.Conn.Exec(ctx,
		`DELETE FROM deck_config WHERE user_id=$1 AND deck_id=$2 RETURNING name`, userId, deckId)
	if err != nil {
		return fmt.Errorf("delete deck failed: %w", err)
	}

	return nil
}

func (d *Storage) UpdateLeitner(ctx context.Context, ul models.UserLeitner) error {
	row := d.Conn.QueryRow(ctx,
		`UPDATE user_leitner SET user_id=$1, card_id=$2, box=$3, cool_down=$4 WHERE leitner_id=$5 RETURNING leitner_id`,
		ul.UserId, ul.FlashcardId, ul.Box, time.Time(ul.CoolDown), ul.Id)

	return row.Scan(&ul.Id)
}

func (d *Storage) DeleteLeitner(ctx context.Context, id models.LeitnerId) error {
	_, err := d.Conn.Exec(ctx, `DELETE FROM user_leitner WHERE leitner_id=$1`, id)
	if err != nil {
		return fmt.Errorf("failed to delete leitner: %w", err)
	}

	return nil
}

func (d *Storage) UpdateDeck(ctx context.Context, deck models.DeckConfig) error {
	_, err := d.Conn.Exec(ctx, `UPDATE deck_config SET user_id=$1, name=$2 
                   WHERE deck_id=$3`, deck.UserId, deck.Name, deck.DeckId)
	if err != nil {
		return fmt.Errorf("update deck failed: %w", err)
	}

	return nil
}

func (d *Storage) UpdateFlashcard(ctx context.Context,
	id models.FlashcardId,
	w models.Word,
	b models.Backside,
	a models.Answer) error {
	_, err := d.Conn.Exec(ctx, `UPDATE flashcard SET word=$1, backside=$2, answer=$3
WHERE card_id=$4`, w, b, a, id)
	if err != nil {
		return fmt.Errorf("update flashcard failed: %w", err)
	}

	return nil
}
