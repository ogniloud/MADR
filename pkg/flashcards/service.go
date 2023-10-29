package flashcards

import (
	"fmt"

	"github.com/ogniloud/madr/pkg/flashcards/cache"
	"github.com/ogniloud/madr/pkg/flashcards/storage"
)

var (
// ErrDatabaseAccess = fmt.Errorf("error when accessing db")
)

// IService must be
type IService interface {
	LoadDecks(id storage.UserId) (storage.Decks, error)
	GetFlashcards(id storage.DeckId) ([]storage.Flashcard, error)

	AppendFlashcard(id storage.DeckId, flashcard storage.Flashcard) error
	DeleteFlashcard(deckId storage.DeckId, flashcardId storage.FlashcardId) error

	CreateNewDeck(userId storage.UserId, cfg storage.DeckConfig, flashcards []storage.Flashcard) error
	DeleteDeck(userId storage.UserId, deckId storage.DeckId) error
	Cache() cache.Cache
}

type Service struct {
	s storage.Storage

	// temporary
	c cache.Cache
}

func NewService(s storage.Storage, c cache.Cache) IService {
	return Service{
		s: s,
		c: c,
	}
}

func (s Service) LoadDecks(id storage.UserId) (storage.Decks, error) {
	if decksAny, ok := s.c.Load(id); ok {
		decks := decksAny.(storage.Decks)
		return decks, nil
	}

	decks, err := s.s.GetDecksByUserId(id)
	if err != nil {
		return nil, err
	}
	go s.c.Store(id, decks)

	return decks, nil
}

func (s Service) GetFlashcards(id storage.DeckId) ([]storage.Flashcard, error) {
	return s.s.GetFlashcardsByDeckId(id)
}

func (s Service) AppendFlashcard(id storage.DeckId, flashcard storage.Flashcard) error {
	return s.s.PutFlashcard(id, flashcard)
}

func (s Service) DeleteFlashcard(deckId storage.DeckId, flashcardId storage.FlashcardId) error {
	return s.s.DeleteFlashcardFromDeck(deckId, flashcardId)
}

func (s Service) CreateNewDeck(userId storage.UserId, cfg storage.DeckConfig, flashcards []storage.Flashcard) error {
	if len(flashcards) == 0 {
		return fmt.Errorf("empty deck")
	}

	decks, err := s.LoadDecks(userId)
	if err != nil {
		return err
	}

	if err := s.s.PutNewDeck(cfg); err != nil {
		return err
	}

	if err := s.s.PutAllFlashcards(cfg.DeckId, flashcards); err != nil {
		return err
	}

	decks[cfg.DeckId] = cfg // maybe critical section!

	return nil
}

func (s Service) DeleteDeck(userId storage.UserId, deckId storage.DeckId) error {
	decks, err := s.LoadDecks(userId)
	if err != nil {
		return err
	}

	if err := s.s.DeleteDeck(deckId); err != nil {
		return err
	}

	delete(decks, deckId)
	return nil
}

func (s Service) Cache() cache.Cache {
	return s.c
}
