package study_test

import (
	"io"
	"slices"
	"strconv"
	"testing"
	"time"

	"github.com/charmbracelet/log"
	"github.com/ogniloud/madr/internal/flashcards/cache"
	"github.com/ogniloud/madr/internal/flashcards/models"
	"github.com/ogniloud/madr/internal/flashcards/services/deck"
	"github.com/ogniloud/madr/internal/flashcards/services/study"
	"github.com/ogniloud/madr/internal/flashcards/storage/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type testingSuite struct {
	suite.Suite
	s study.IStudyService
	m *mocks.Storage
}

func (t *testingSuite) SetupTest() {
	t.m = &mocks.Storage{}
	serv := deck.NewService(t.m, cache.New(), log.New(io.Discard))
	t.s = study.NewService(serv, 5)
}

var (
	flashcard       = []models.FlashcardId{1, 2, 3, 4, 5}
	flashcardCalled = make([]bool, 5)
	leitners        = []models.UserLeitner{
		{
			Id:          1,
			UserId:      1,
			FlashcardId: 1,
			Box:         0,
			CoolDown:    models.CoolDown{State: time.Time{}},
		},
		{
			Id:          2,
			UserId:      1,
			FlashcardId: 2,
			Box:         0,
			CoolDown:    models.CoolDown{},
		},
		{
			Id:          3,
			UserId:      1,
			FlashcardId: 3,
			Box:         1,
			CoolDown:    models.CoolDown{},
		},
		{
			Id:          4,
			UserId:      1,
			FlashcardId: 4,
			Box:         2,
			CoolDown:    models.CoolDown{},
		},
		{
			Id:          5,
			UserId:      1,
			FlashcardId: 5,
			Box:         2,
			CoolDown:    models.CoolDown{},
		},
	}
)

func (t *testingSuite) Test_GetNextRandomDeck() {
	t.m.On("GetUserInfo", 1).
		Return(models.UserInfo{
			UserId: 1,
			MaxBox: 3,
		}, nil)
	t.m.On("GetFlashcardsIdByDeckId", 1).
		Return(flashcard, nil)
	t.m.On("UpdateLeitner", mock.Anything).Return(nil)
	t.m.On("GetLeitnerByUserIdCardId", 1, mock.AnythingOfType("int")).
		Return(func(i, j int) (models.UserLeitner, error) {
			if flashcardCalled[j-1] {
				return models.UserLeitner{
					Id:          0,
					UserId:      0,
					FlashcardId: 0,
					Box:         0,
					CoolDown:    models.CoolDown{State: time.Time{}.Add(time.Hour)},
				}, nil
			}
			return leitners[j-1], nil
		})

	cards := make([]models.FlashcardId, 0, 5)
	for i := 0; i < len(flashcard); i++ {
		t.Run(strconv.Itoa(i), func() {
			id, err := t.s.GetNextRandomDeck(1, 1, models.CoolDown{State: time.Time{}.Add(time.Second)})
			if assert.NoError(t.T(), err) {
				cards = append(cards, id)
				flashcardCalled[id-1] = true
			}
		})
	}
	_, err := t.s.GetNextRandomDeck(1, 1, models.CoolDown{State: time.Time{}.Add(time.Second)})
	assert.Error(t.T(), err)

	slices.Sort(cards)
	assert.Equal(t.T(), flashcard, cards)

	_, err = t.s.GetNextRandomDeck(1, 1, models.CoolDown{State: time.Time{}.Add(2 * time.Hour)})
	assert.NoError(t.T(), err)
}

var leitners2 = []models.UserLeitner{
	{
		Id:          1,
		UserId:      1,
		FlashcardId: 1,
		Box:         0,
		CoolDown:    models.CoolDown{State: time.Time{}},
	},
	{
		Id:          2,
		UserId:      1,
		FlashcardId: 2,
		Box:         0,
		CoolDown:    models.CoolDown{},
	},
	{
		Id:          3,
		UserId:      1,
		FlashcardId: 3,
		Box:         1,
		CoolDown:    models.CoolDown{},
	},
	{
		Id:          4,
		UserId:      1,
		FlashcardId: 4,
		Box:         2,
		CoolDown:    models.CoolDown{},
	},
	{
		Id:          5,
		UserId:      1,
		FlashcardId: 5,
		Box:         2,
		CoolDown:    models.CoolDown{},
	},
	{
		Id:          6,
		UserId:      1,
		FlashcardId: 6,
		Box:         1,
		CoolDown:    models.CoolDown{},
	},
	{
		Id:          7,
		UserId:      1,
		FlashcardId: 7,
		Box:         2,
		CoolDown:    models.CoolDown{},
	},
	{
		Id:          8,
		UserId:      1,
		FlashcardId: 8,
		Box:         2,
		CoolDown:    models.CoolDown{},
	},
}

var (
	flashcard2       = []models.FlashcardId{6, 7, 8}
	flashcard2Called = [8]bool{}
)

const length = 8

func (t *testingSuite) Test_GetNextRandom() {
	t.m.On("GetDecksByUserId", 1).Return(models.Decks{1: {}, 2: {}}, nil)
	t.m.On("GetUserInfo", 1).
		Return(models.UserInfo{
			UserId: 1,
			MaxBox: 3,
		}, nil)
	t.m.On("GetFlashcardsIdByDeckId", 1).
		Return(flashcard, nil)
	t.m.On("GetFlashcardsIdByDeckId", 2).
		Return(flashcard2, nil)
	t.m.On("UpdateLeitner", mock.Anything).Return(nil)
	t.m.On("GetLeitnerByUserIdCardId", 1, mock.AnythingOfType("int")).
		Return(func(i, j int) (models.UserLeitner, error) {
			if flashcard2Called[j-1] {
				return models.UserLeitner{
					Id:          0,
					UserId:      0,
					FlashcardId: 0,
					Box:         0,
					CoolDown:    models.CoolDown{State: time.Time{}.Add(time.Hour)},
				}, nil
			}
			return leitners2[j-1], nil
		})

	cards := make([]models.FlashcardId, 0, length)
	for i := 0; i < length; i++ {
		t.Run(strconv.Itoa(i), func() {
			id, err := t.s.GetNextRandom(1, models.CoolDown{State: time.Time{}.Add(time.Second)})
			if assert.NoError(t.T(), err) {
				cards = append(cards, id)
				flashcard2Called[id-1] = true
			}
		})
	}
	_, err := t.s.GetNextRandom(1, models.CoolDown{State: time.Time{}.Add(time.Second)})
	assert.Error(t.T(), err)

	slices.Sort(cards)
	assert.Equal(t.T(), append(flashcard, flashcard2...), cards)

	_, err = t.s.GetNextRandom(1, models.CoolDown{State: time.Time{}.Add(2 * time.Hour)})
	assert.NoError(t.T(), err)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(testingSuite))
}
