package deck

import (
	"fmt"

	"github.com/ogniloud/madr/internal/flashcards/cache"
	"github.com/ogniloud/madr/internal/flashcards/models"
	"github.com/ogniloud/madr/internal/flashcards/storage"

	"github.com/charmbracelet/log"
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
	c      cache.Cache
	logger *log.Logger
}

// NewService creates a new IService creature.
func NewService(s storage.Storage, c cache.Cache, logger *log.Logger) IService {
	return &Service{
		Storage: s,
		c:       c,
		logger:  logger,
	}
}

// LoadDecks loads user's decks in memory if needed and returns the decks.
func (s *Service) LoadDecks(id models.UserId) (models.Decks, error) {
	if decksAny, err := s.c.Load(id); err == nil {
		if decks, ok := decksAny.(models.Decks); ok {
			return decks, nil // returned from cache
		}
	}

	decks, err := s.GetDecksByUserId(id)
	if err != nil {
		return nil, err
	}

	if err := s.Cache().Store(id, decks); err != nil {
		s.logger.Errorf("cache store failed: %v", err)
	}

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
	if err := s.Cache().Store(userId, decks); err != nil {
		s.logger.Errorf("cache store failed: %v", err)
		_ = s.Cache().Delete(userId) // non-consistent cache
	}

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
	if err := s.Cache().Store(userId, decks); err != nil {
		s.logger.Errorf("cache store failed: %v", err)
		_ = s.Cache().Delete(userId) // non-consistent cache
	}

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
