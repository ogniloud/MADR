package deck

import (
	"context"
	"fmt"
	"time"

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
	LoadDecks(ctx context.Context, id models.UserId) (models.Decks, error)

	// NewDeckWithFlashcards creates a new deck and puts all the flashcards there.
	// Returns id of the new deck if there's no error.
	NewDeckWithFlashcards(ctx context.Context, userId models.UserId, cfg models.DeckConfig, flashcards []models.Flashcard) (models.DeckId, error)

	// DeleteDeck deletes the whole deck from user. It doesn't remove flashcards from the deck.
	DeleteDeck(ctx context.Context, userId models.UserId, deckId models.DeckId) error

	// UserMaxBox returns the maximum amount of boxes from the user.
	UserMaxBox(ctx context.Context, uid models.UserId) (models.Box, error)
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
func (s *Service) LoadDecks(ctx context.Context, id models.UserId) (_ models.Decks, err error) {
	if decksAny, err := s.c.Load(id); err == nil {
		if decks, ok := decksAny.(models.Decks); ok {
			return decks, nil // returned from cache
		}
	}

	decks, err := s.GetDecksByUserId(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := s.Cache().Store(id, decks); err != nil {
		s.logger.Errorf("cache store failed: %v", err)
	}

	return decks, nil
}

// NewDeckWithFlashcards creates a new deck and puts all the flashcards there.
func (s *Service) NewDeckWithFlashcards(ctx context.Context, userId models.UserId, cfg models.DeckConfig, flashcards []models.Flashcard) (_ models.DeckId, err error) {
	if len(flashcards) == 0 {
		return 0, ErrEmptyFlashcardsSlice
	}

	decks, err := s.LoadDecks(ctx, userId)
	if err != nil {
		return 0, err
	}

	var id models.DeckId
	if id, err = s.PutNewDeck(ctx, cfg); err != nil {
		return 0, err
	}

	var ids []models.FlashcardId
	if ids, err = s.PutAllFlashcards(ctx, id, flashcards); err != nil {
		return 0, err
	}

	uls := make([]models.UserLeitner, len(flashcards))
	for i := 0; i < len(flashcards); i++ {
		uls[i] = models.UserLeitner{
			UserId:      userId,
			FlashcardId: ids[i],
			Box:         0,
			CoolDown:    models.CoolDown(time.Now()),
		}
	}
	if _, err = s.PutAllUserLeitner(ctx, uls); err != nil {
		return 0, err
	}

	decks[cfg.DeckId] = cfg // maybe critical section!
	if err = s.Cache().Store(userId, decks); err != nil {
		s.logger.Errorf("cache store failed: %v", err)
		_ = s.Cache().Delete(userId) // non-consistent cache
	}

	return id, nil
}

// DeleteDeck deletes the whole deck from user. It doesn't remove flashcards from the deck.
func (s *Service) DeleteDeck(ctx context.Context, userId models.UserId, deckId models.DeckId) (err error) {
	decks, err := s.LoadDecks(ctx, userId)
	if err != nil {
		return err
	}

	if err := s.DeleteUserDeck(ctx, userId, deckId); err != nil {
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
func (s *Service) UserMaxBox(ctx context.Context, uid models.UserId) (_ models.Box, err error) {
	info, err := s.GetUserInfo(ctx, uid)
	if err != nil {
		return 0, err
	}

	return info.MaxBox, nil
}

func (s *Service) Cache() cache.Cache {
	return s.c
}
