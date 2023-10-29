package flashcards_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/ogniloud/madr/pkg/flashcards"
	"github.com/ogniloud/madr/pkg/flashcards/storage"
	"github.com/ogniloud/madr/pkg/flashcards/storage/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LeitnerSuite struct {
	suite.Suite
	s flashcards.IService
}

func (l *LeitnerSuite) SetupTest() {
	s := &mocks.Storage{}
	l.s = flashcards.NewService(s, &sync.Map{})

	s.On("GetDecksByUserId", 1).
		Return(user1, nil)
	s.On("GetDecksByUserId", 666).
		Return(nil, fmt.Errorf("user not found"))
}

func (l *LeitnerSuite) Test_LoadDecks() {
	decks, err := l.s.LoadDecks(1)
	assert.Nil(l.T(), err)
	assert.Equal(l.T(), user1, decks)

	_, err = l.s.LoadDecks(666)
	assert.Error(l.T(), err)

	l.Run("cache check", func() {
		var cachedDecks storage.Decks

		cd, ok := l.s.Cache().Load(1)
		assert.True(l.T(), ok)
		cachedDecks = cd.(storage.Decks)

		assert.Equal(l.T(), cachedDecks, user1)
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(LeitnerSuite))
}

var user1 = storage.Decks{
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
