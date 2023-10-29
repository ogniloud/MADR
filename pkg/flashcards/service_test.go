package flashcards_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/ogniloud/madr/pkg/flashcards"
	"github.com/ogniloud/madr/pkg/flashcards/storage"
	"github.com/ogniloud/madr/pkg/flashcards/storage/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type LeitnerSuite struct {
	suite.Suite
	s        flashcards.Service
	st       *mocks.Storage
	userData storage.Decks
}

func (l *LeitnerSuite) SetupTest() {
	l.userData = storage.Decks{
		1: storage.DeckConfig{
			DeckId: 1,
			UserId: 1,
			Name:   "Deck1",
		},
		2: storage.DeckConfig{
			DeckId: 2,
			UserId: 1,
			Name:   "Deck2",
		},
		4: storage.DeckConfig{
			DeckId: 4,
			UserId: 1,
			Name:   "Deck4",
		},
	}

	// s := mocks.NewStorage(l.T())
	s := &mocks.Storage{}
	l.st = s
	l.s = flashcards.NewService(s, &sync.Map{})

	s.On("GetDecksByUserId", 1).
		Return(l.userData, nil).Once()
	s.On("GetDecksByUserId", 1).
		Return(nil, fmt.Errorf("db problem"))
}

func (l *LeitnerSuite) Test_LoadDecks() {
	s := l.st

	s.On("GetDecksByUserId", 2).
		Return(l.userData, nil).Once()
	s.On("GetDecksByUserId", 666).
		Return(nil, fmt.Errorf("user not found")).Once()

	l.Run("load #1", func() {
		decks, err := l.s.LoadDecks(1)
		if assert.NoError(l.T(), err) {
			assert.Equal(l.T(), l.userData, decks)
		}

		l.Run("cache check", func() {
			var cachedDecks storage.Decks

			cd, ok := l.s.Cache().Load(1)
			if assert.True(l.T(), ok) {
				cachedDecks = cd.(storage.Decks)
			}

			assert.Equal(l.T(), cachedDecks, l.userData)
		})
	})

	l.Run("load #2", func() {
		decks, err := l.s.LoadDecks(2)
		if assert.NoError(l.T(), err) {
			assert.Equal(l.T(), l.userData, decks)
		}
	})

	l.Run("not found", func() {
		_, err := l.s.LoadDecks(666)
		assert.Error(l.T(), err)
	})

	// take from cache, must not panic
	l.Run("pulling from cache", func() {
		_, err := l.st.GetDecksByUserId(1)
		assert.Error(l.T(), err)

		decks, err := l.s.LoadDecks(1)
		if assert.NoError(l.T(), err) {
			assert.Equal(l.T(), l.userData, decks)
		}
	})
}

func (l *LeitnerSuite) Test_CreateNewDeck() {
	s := l.st
	s.On("PutNewDeck", mock.AnythingOfType("storage.DeckConfig")).
		Return(nil).Once()
	s.On("PutAllFlashcards",
		mock.AnythingOfType("int"),
		mock.AnythingOfType("[]storage.Flashcard")).
		Return(nil).Once()

	s.On("GetDecksByUserId", 1).
		Return(l.userData, nil).Once()

	cfg := storage.DeckConfig{
		DeckId: 10,
		UserId: 1,
		Name:   "Deck10",
	}
	l.userData[10] = cfg

	cards := []storage.Flashcard{
		{
			Id:     1,
			W:      "Aboba",
			B:      storage.Backside{},
			DeckId: 10,
		},
		{
			Id:     2,
			W:      "Durdom",
			B:      storage.Backside{},
			DeckId: 10,
		},
	}
	assert.Error(l.T(), l.s.CreateNewDeck(1, storage.DeckConfig{}, []storage.Flashcard{}))

	assert.Nil(l.T(), l.s.CreateNewDeck(1, cfg, cards))

	decks, err := l.s.LoadDecks(1)
	if assert.NoError(l.T(), err) {
		assert.Equal(l.T(), l.userData, decks)
	}
}

func (l *LeitnerSuite) Test_DeleteDeck() {
	s := l.st
	s.On("DeleteDeck", 3).
		Return(fmt.Errorf("deck not found")).Once()
	s.On("DeleteDeck", mock.AnythingOfType("int")).
		Return(nil).Once()
	s.On("GetDecksByUserId", 3).
		Return(nil, fmt.Errorf("user not found"))

	l.Run("deck not found", func() {
		assert.Error(l.T(), l.s.DeleteDeck(1, 3))
	})

	l.Run("user not found", func() {
		assert.Error(l.T(), l.s.DeleteDeck(3, 1))
	})

	delete(l.userData, 4)
	l.Run("success delete", func() {
		assert.NoError(l.T(), l.s.DeleteDeck(1, 4))
		decks, err := l.s.LoadDecks(1)
		if assert.NoError(l.T(), err) {
			assert.Equal(l.T(), l.userData, decks)
		}
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(LeitnerSuite))
}