package flashcards

import "github.com/ogniloud/madr/pkg/flashcards/storage"

type DeckConfig struct {
	Name string `json:"name"`
}

// Service must be
type Service interface {
	LoadDecks() ([]storage.DeckId, error)
	GetFlashcards(id storage.DeckId) ([]storage.Flashcard, error)

	/* decks can be changed only if they are not public */
	AppendFlashcard(id storage.DeckId, flashcard storage.Flashcard) error
	DeleteFlashcard(deckId storage.DeckId, flashcardId storage.FlashcardId) error

	CreateNewDeck(cfg DeckConfig, flashcards []storage.Flashcard) error
	DeleteDeck(id storage.DeckId) error
}
