package study_test

import (
	"sync"
	"testing"

	"github.com/ogniloud/madr/pkg/flashcards/services/deck"
	"github.com/ogniloud/madr/pkg/flashcards/services/study"
	"github.com/ogniloud/madr/pkg/flashcards/storage/mocks"
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

func (t *testingSuite) Test_GetNextRandom() {

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(testingSuite))
}
