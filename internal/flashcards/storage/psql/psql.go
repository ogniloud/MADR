package psql

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

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
	row := d.Conn.QueryRow(context.Background(), `SELECT * FROM flashcard WHERE card_id=$1`, id)

	flashcard := models.Flashcard{}
	strBackside := ""

	err := row.Scan(&flashcard.Id, &flashcard.W, &strBackside, &flashcard.DeckId)
	if err != nil {
		return models.Flashcard{}, fmt.Errorf("psql error: %w", err)
	}

	err = json.NewDecoder(strings.NewReader(strBackside)).Decode(&flashcard.B)
	if err != nil {
		return models.Flashcard{}, fmt.Errorf("json error: %w", err)
	}

	return flashcard, nil
}

func (d *DeckStorage) GetLeitnerByUserIdCardId(id models.UserId, flashcardId models.FlashcardId) (models.UserLeitner, error) {
	row := d.Conn.QueryRow(context.Background(), `SELECT * FROM user_leitner WHERE user_id=$1 AND card_id=$2`, id, flashcardId)

	l := models.UserLeitner{}

	err := row.Scan(&l.Id, &l.UserId, &l.FlashcardId, &l.Box, &l.CoolDown.State)
	if err != nil {
		return models.UserLeitner{}, fmt.Errorf("psql error: %w", err)
	}

	return l, nil
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
