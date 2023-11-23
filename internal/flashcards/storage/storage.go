package storage

import (
	"context"

	"github.com/ogniloud/madr/internal/flashcards/models"
)

// Storage is an interface for accessing to database.
//
//go:generate go run github.com/vektra/mockery/v2@v2.36.0 --name Storage
type Storage interface {
	// GetDecksByUserId returns all decks the user has.
	GetDecksByUserId(ctx context.Context, id models.UserId) (models.Decks, error)

	// GetFlashcardsIdByDeckId returns the id of the flashcards in the deck.
	GetFlashcardsIdByDeckId(ctx context.Context, id models.DeckId) ([]models.FlashcardId, error)

	// GetFlashcardById returns models.Flashcard by its id.
	GetFlashcardById(ctx context.Context, id models.FlashcardId) (models.Flashcard, error)

	// GetLeitnerByUserIdCardId returns models.UserLeitner by user and card.
	GetLeitnerByUserIdCardId(ctx context.Context, id models.UserId, flashcardId models.FlashcardId) (models.UserLeitner, error)

	// GetUserInfo returns models.UserInfo by models.UserId.
	GetUserInfo(ctx context.Context, uid models.UserId) (models.UserInfo, error)

	// PutAllFlashcards appends flashcards to the deck by appending new rows in models.Flashcard storage.
	PutAllFlashcards(ctx context.Context, id models.DeckId, cards []models.Flashcard) ([]models.FlashcardId, error)

	// PutNewDeck appends a new deck configuration.
	PutNewDeck(ctx context.Context, config models.DeckConfig) (models.DeckId, error)

	// PutAllUserLeitner appends all the models.UserLeitner in storage.
	// Returns a slice of ids in order of the slice of models.UserLeitner.
	PutAllUserLeitner(ctx context.Context, uls []models.UserLeitner) ([]models.LeitnerId, error)

	// DeleteFlashcardFromDeck deletes a record from models.Flashcard storage about flashcard.
	DeleteFlashcardFromDeck(ctx context.Context, cardId models.FlashcardId) error

	// DeleteUserDeck removes the entry from models.DeckConfig that user with userId has a deck.
	DeleteUserDeck(ctx context.Context, userId models.UserId, id models.DeckId) error

	// UpdateLeitner updates a record in the database when the models.LeitnerId matches.
	UpdateLeitner(ctx context.Context, ul models.UserLeitner) error
}
