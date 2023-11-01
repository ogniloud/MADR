package study_test

import (
	"sync"
	"testing"

	"github.com/ogniloud/madr/pkg/flashcards/models"
	"github.com/ogniloud/madr/pkg/flashcards/services/deck"
	"github.com/ogniloud/madr/pkg/flashcards/services/study"
	"github.com/ogniloud/madr/pkg/flashcards/storage/mocks"
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

var flashcards = []models.Flashcard{
	{
		Id: 1,
		W:  "Aboba1",
		B: models.Backside{
			1, "Абоба1",
		},
		DeckId: 2,
	},
	{
		Id: 2,
		W:  "Aboba2",
		B: models.Backside{
			1, "Абоба2",
		},
		DeckId: 2,
	},
	{
		Id: 3,
		W:  "Aboba3",
		B: models.Backside{
			1, "Абоба3",
		},
		DeckId: 5,
	},
	{
		Id: 4,
		W:  "Aboba4",
		B: models.Backside{
			1, "Абоба4",
		},
		DeckId: 3,
	},
}

func (t *testingSuite) Test_GetNRandom() {
	t.m.On("GetCardsByUserCDBox", mock.Anything, mock.Anything, mock.Anything).
		Return(flashcards, nil).Once()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(testingSuite))
}
