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
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(LeitnerSuite))
}

var user1 = storage.Decks{
	1: storage.DeckConfig{
		Id:   1,
		Name: "Deck1",
	},
	2: storage.DeckConfig{
		Id:   2,
		Name: "Deck2",
	},
	4: storage.DeckConfig{
		Id:   4,
		Name: "Deck4",
	},
}
