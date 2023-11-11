package psql

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/ogniloud/madr/internal/db"
	"github.com/ogniloud/madr/internal/flashcards/models"
)

type DeckStorage struct {
	Conn *db.PSQLDatabase
}

func (d *DeckStorage) GetDecksByUserId(id models.UserId) (models.Decks, error) {
	rows, err := d.Conn.Query(context.Background(), `SELECT * FROM deck_config WHERE user_id=$1`, id)
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

func (d *DeckStorage) GetFlashcardsIdByDeckId(id models.DeckId) ([]models.FlashcardId, error) {
	rows, err := d.Conn.Query(context.Background(), `SELECT card_id FROM flashcard WHERE deck_id=$1`, id)
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

func (d *DeckStorage) GetFlashcardById(id models.FlashcardId) (models.Flashcard, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DeckStorage) GetLeitnerByUserIdCardId(id models.UserId, flashcardId models.FlashcardId) (models.UserLeitner, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DeckStorage) GetUserInfo(uid models.UserId) (models.UserInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DeckStorage) PutAllFlashcards(id models.DeckId, cards []models.Flashcard) ([]models.FlashcardId, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DeckStorage) PutNewDeck(config models.DeckConfig) (models.DeckId, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DeckStorage) PutAllUserLeitner(uls []models.UserLeitner) ([]models.LeitnerId, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DeckStorage) DeleteFlashcardFromDeck(cardId models.FlashcardId) error {
	//TODO implement me
	panic("implement me")
}

func (d *DeckStorage) DeleteUserDeck(userId models.UserId, id models.DeckId) error {
	//TODO implement me
	panic("implement me")
}

func (d *DeckStorage) UpdateLeitner(ul models.UserLeitner) error {
	//TODO implement me
	panic("implement me")
}
