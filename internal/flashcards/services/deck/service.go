package deck

import (
	"fmt"

	"github.com/ogniloud/madr/internal/flashcards/cache"
	"github.com/ogniloud/madr/internal/flashcards/models"
	"github.com/ogniloud/madr/internal/flashcards/storage"
)

var ErrEmptyFlashcardsSlice = fmt.Errorf("empty flashcards slice")

// IService is a service for configuration user decks and flashcards.
type IService interface {
	// Storage is embedded to IService for accessing to db methods.
	// temporary solution.
	storage.Storage

	// LoadDecks loads user's decks in memory if needed and returns the decks.
	LoadDecks(id models.UserId) (models.Decks, error)

	// NewDeckWithFlashcards creates a new deck and puts all the flashcards there.
	NewDeckWithFlashcards(userId models.UserId, cfg models.DeckConfig, flashcards []models.Flashcard) error

	// DeleteDeck deletes the whole deck from user. It doesn't remove flashcards from the deck.
	DeleteDeck(userId models.UserId, deckId models.DeckId) error

	// UserMaxBox returns the maximum amount of boxes from the user.
	UserMaxBox(uid models.UserId) (models.Box, error)
	Cache() cache.Cache
}

// Service is an IService implementation.
type Service struct {
	storage.Storage

	c cache.Cache
}

// NewService creates a new IService creature.
func NewService(s storage.Storage, c cache.Cache) IService {
	return &Service{
		Storage: s,
		c:       c,
	}
}

// LoadDecks loads user's decks in memory if needed and returns the decks.
func (s *Service) LoadDecks(id models.UserId) (models.Decks, error) {
	if decksAny, ok := s.c.Load(id); ok {
		decks := decksAny.(models.Decks)
		return decks, nil
	}

	decks, err := s.GetDecksByUserId(id)
	if err != nil {
		return nil, err
	}

	s.c.Store(id, decks)

	return decks, nil
}

// NewDeckWithFlashcards creates a new deck and puts all the flashcards there.
func (s *Service) NewDeckWithFlashcards(userId models.UserId, cfg models.DeckConfig, flashcards []models.Flashcard) error {
	if len(flashcards) == 0 {
		return ErrEmptyFlashcardsSlice
	}

	decks, err := s.LoadDecks(userId)
	if err != nil {
		return err
	}

	if err := s.PutNewDeck(cfg); err != nil {
		return err
	}

	if err := s.PutAllFlashcards(cfg.DeckId, flashcards); err != nil {
		return err
	}

	decks[cfg.DeckId] = cfg // maybe critical section!
	s.Cache().Store(userId, decks)

	return nil
}

// DeleteDeck deletes the whole deck from user. It doesn't remove flashcards from the deck.
func (s *Service) DeleteDeck(userId models.UserId, deckId models.DeckId) error {
	decks, err := s.LoadDecks(userId)
	if err != nil {
		return err
	}

	if err := s.DeleteUserDeck(userId, deckId); err != nil {
		return err
	}

	delete(decks, deckId)
	s.Cache().Store(userId, decks)

	return nil
}

// UserMaxBox returns the maximum amount of boxes from the user.
func (s *Service) UserMaxBox(uid models.UserId) (models.Box, error) {
	info, err := s.GetUserInfo(uid)
	if err != nil {
		return 0, err
	}

	return info.MaxBox, nil
}

func (s *Service) Cache() cache.Cache {
	return s.c
}
