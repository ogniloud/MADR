package deck

import (
	"fmt"

	"github.com/ogniloud/madr/pkg/flashcards/cache"
	"github.com/ogniloud/madr/pkg/flashcards/models"
	"github.com/ogniloud/madr/pkg/flashcards/storage"
)

var ErrEmptyFlashcardsSlice = fmt.Errorf("empty flashcards slice")

type Service struct {
	storage.Storage

	// temporary
	c cache.Cache
}

func NewService(s storage.Storage, c cache.Cache) Service {
	return Service{
		Storage: s,
		c:       c,
	}
}

// LoadDecks загружает колоды пользователя в кэш и возвращает.
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

// CreateNewDeck Создаёт новую колоду и вставляет туда карточки.
func (s *Service) CreateNewDeck(userId models.UserId, cfg models.DeckConfig, flashcards []models.Flashcard) error {
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

// DeleteDeck удаляет колоду у пользователя.
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

// UserMaxBox возвращает максимальное количество Box'ов у пользователя.
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
