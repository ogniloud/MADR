package storage

import (
	"github.com/ogniloud/madr/internal/flashcards/models"
)

// Storage is an interface for accessing to database.
//
//go:generate go run github.com/vektra/mockery/v2@v2.36.0 --name Storage
type Storage interface {
	// GetDecksByUserId returns all decks the user has.
	GetDecksByUserId(id models.UserId) (models.Decks, error)

	// GetFlashcardsIdByDeckId returns the id of the flashcards in the deck.
	GetFlashcardsIdByDeckId(id models.DeckId) ([]models.FlashcardId, error)

	// GetFlashcardById returns models.Flashcard by its id.
	GetFlashcardById(id models.FlashcardId) (models.Flashcard, error)

	// GetLeitnerByUserIdCardId returns models.UserLeitner by user and card.
	GetLeitnerByUserIdCardId(id models.UserId, flashcardId models.FlashcardId) (models.UserLeitner, error)

	// GetUserInfo returns models.UserInfo by models.UserId.
	GetUserInfo(uid models.UserId) (models.UserInfo, error)

	// PutAllFlashcards appends flashcards to the deck by appending new rows in models.Flashcard storage.
	PutAllFlashcards(id models.DeckId, cards []models.Flashcard) error

	// PutNewDeck appends a new deck configuration.
	PutNewDeck(config models.DeckConfig) error

	// PutAllUserLeitner appends all the models.UserLeitner in storage.
	PutAllUserLeitner(uls []models.UserLeitner) error

	// DeleteFlashcardFromDeck deletes a record from models.Flashcard storage about flashcard.
	DeleteFlashcardFromDeck(cardId models.FlashcardId) error

	// DeleteUserDeck removes the entry from models.DeckConfig that user with userId has a deck.
	DeleteUserDeck(userId models.UserId, id models.DeckId) error

	// UpdateLeitner updates a record in the database when the models.LeitnerId matches.
	UpdateLeitner(ul models.UserLeitner) error
}
