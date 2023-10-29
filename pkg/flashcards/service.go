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

	GetRandom(id storage.UserId, cd storage.CoolDown, limits map[storage.Box]int) ([]storage.FlashcardId, error)
	GetRandomDeck(id storage.UserId, cd storage.CoolDown, deckId storage.DeckId, limits map[storage.Box]int) ([]storage.FlashcardId, error)

	GetUserMaxBox(uid storage.UserId) (storage.Box, error)
	Cache() cache.Cache
}

type Service struct {
	S storage.Storage

	// temporary
	c cache.Cache
}

func NewService(s storage.Storage, c cache.Cache) IService {
	return Service{
		S: s,
		c: c,
	}
}

func (s Service) LoadDecks(id storage.UserId) (storage.Decks, error) {
	if decksAny, ok := s.c.Load(id); ok {
		decks := decksAny.(storage.Decks)
		return decks, nil
	}

	decks, err := s.S.GetDecksByUserId(id)
	if err != nil {
		return nil, err
	}
	s.c.Store(id, decks)

	return decks, nil
}

func (s Service) GetFlashcards(id storage.DeckId) ([]storage.Flashcard, error) {
	return s.S.GetFlashcardsByDeckId(id)
}

func (s Service) AppendFlashcard(id storage.DeckId, flashcard storage.Flashcard) error {
	return s.S.PutAllFlashcards(id, []storage.Flashcard{flashcard})
}

func (s Service) DeleteFlashcard(deckId storage.DeckId, flashcardId storage.FlashcardId) error {
	return s.S.DeleteFlashcardFromDeck(deckId, flashcardId)
}

func (s Service) CreateNewDeck(userId storage.UserId, cfg storage.DeckConfig, flashcards []storage.Flashcard) error {
	if len(flashcards) == 0 {
		return fmt.Errorf("empty deck")
	}

	decks, err := s.LoadDecks(userId)
	if err != nil {
		return err
	}

	if err := s.S.PutNewDeck(cfg); err != nil {
		return err
	}

	if err := s.S.PutAllFlashcards(cfg.DeckId, flashcards); err != nil {
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

	if err := s.S.DeleteDeck(deckId); err != nil {
		return err
	}

	delete(decks, deckId)
	return nil
}

func (s Service) GetRandom(id storage.UserId, cd storage.CoolDown, limits map[storage.Box]int) ([]storage.FlashcardId, error) {
	return s.S.GetCardsByUserCDBox(id, cd, limits)
}

func (s Service) GetRandomDeck(id storage.UserId, cd storage.CoolDown, deckId storage.DeckId, limits map[storage.Box]int) ([]storage.FlashcardId, error) {
	return s.S.GetCardsByUserCDBoxDeck(id, cd, limits, deckId)
}

func (s Service) GetUserMaxBox(uid storage.UserId) (storage.Box, error) {
	info, err := s.S.GetUserInfo(uid)
	if err != nil {
		return 0, err
	}

	return info.MaxBox, nil
}

func (s Service) Cache() cache.Cache {
	return s.c
}
