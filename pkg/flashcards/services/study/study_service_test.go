package study_test

import (
	"sync"
	"testing"

	"github.com/ogniloud/madr/pkg/flashcards/models"
	"github.com/ogniloud/madr/pkg/flashcards/services/deck"
	"github.com/ogniloud/madr/pkg/flashcards/services/study"
	"github.com/ogniloud/madr/pkg/flashcards/storage/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type testingSuite struct {
	suite.Suite
	s *study.StudyService
	m *mocks.Storage
}

func (t *testingSuite) SetupTest() {
	t.m = &mocks.Storage{}
	serv := deck.NewService(t.m, &sync.Map{})
	t.s = study.NewStudy(&serv, 5)
}

var ids = map[models.FlashcardId]struct{}{
	1: {}, 2: {}, 3: {}, 4: {},
}
var flashcards = []models.FlashcardId{1, 2, 3, 4}

func (t *testingSuite) Test_GetNRandom() {
	t.m.On("GetCardsByUserCDBox", mock.Anything, mock.Anything, mock.Anything).
		Return(nil, nil).Once()

	var cards map[models.FlashcardId]struct{}

	t.Run("getting cards", func() {
		for i := 0; i < 4; i++ {
			card, err := t.s.GetRandom(1, 1)
			assert.NoError(t.T(), err)
			cards[card] = struct{}{}
		}
	})

	t.Run("empty cache", func() {
		_, err := t.s.GetRandom(1, 1)
		assert.Error(t.T(), err)
	})

	t.Run("equal", func() {
		assert.Equal(t.T(), cards, ids)
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(testingSuite))
}
