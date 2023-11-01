package deck

import (
	"fmt"
	"slices"

	"github.com/ogniloud/madr/pkg/flashcards/cache"
	"github.com/ogniloud/madr/pkg/flashcards/models"
	"github.com/ogniloud/madr/pkg/flashcards/storage"
)

var (
	ErrEmptyFlashcardsSlice = fmt.Errorf("empty flashcards slice")

	keyRandomCardFmt     = "k%d"
	keyRandomCardDeckFmt = "%dk%d"
)

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

	return nil
}

func (s *Service) RemoveDeckFromUser(userId models.UserId, deckId models.DeckId) error {
	decks, err := s.LoadDecks(userId)
	if err != nil {
		return err
	}

	if err := s.DeleteUserDeck(userId, deckId); err != nil {
		return err
	}

	delete(decks, deckId)
	return nil
}

func (s *Service) Random(id models.UserId, cd models.CoolDown, limits map[models.Box]int) error {
	cards, err := s.GetCardsByUserCDBox(id, cd, limits)
	if err != nil {
		return err
	}

	if len(cards) == 0 {
		return ErrEmptyFlashcardsSlice
	}

	cachedCardsTmp, _ := s.c.Load(fmt.Sprintf(keyRandomCardFmt, id))
	cachedCards := cachedCardsTmp.(cache.CachedRandom)

	cachedCards = append(cachedCards, cards...)
	cachedCards = slices.Clip(cachedCards)

	s.c.Store(fmt.Sprintf(keyRandomCardFmt, id), cachedCards)

	return nil
}

func (s *Service) RandomDeck(id models.UserId, cd models.CoolDown, deckId models.DeckId, limits map[models.Box]int) error {
	cards, err := s.GetCardsByUserCDBoxDeck(id, cd, limits, deckId)
	if err != nil {
		return err
	}

	if len(cards) == 0 {
		return ErrEmptyFlashcardsSlice
	}

	cachedCardsTmp, _ := s.c.Load(fmt.Sprintf(keyRandomCardDeckFmt, deckId, id))
	cachedCards := cachedCardsTmp.(cache.CachedRandom)

	cachedCards = append(cachedCards, cards...)
	cachedCards = slices.Clip(cachedCards)

	s.c.Store(fmt.Sprintf(keyRandomCardDeckFmt, deckId, id), cachedCards)

	return nil
}

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
