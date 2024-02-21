package study_test

import (
	"context"
	"io"
	"slices"
	"strconv"
	"testing"
	"time"

	"github.com/ogniloud/madr/internal/flashcards/cache"
	"github.com/ogniloud/madr/internal/flashcards/models"
	"github.com/ogniloud/madr/internal/flashcards/services/deck"
	"github.com/ogniloud/madr/internal/flashcards/services/study"
	"github.com/ogniloud/madr/internal/flashcards/storage/mocks"

	"github.com/charmbracelet/log"
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
	t.s = study.NewService(serv)
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
			CoolDown:    models.CoolDown{},
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
	flashcardContents = []models.Flashcard{
		{
			Id: 1,
			W:  "word1",
			A:  "answer1",
			B: models.Backside{
				Type:  1,
				Value: "backside1",
			},
			DeckId: 1,
		},
		{
			Id: 2,
			W:  "word2",
			A:  "answer2",
			B: models.Backside{
				Type:  1,
				Value: "backside2",
			},
			DeckId: 1,
		},
		{
			Id: 3,
			W:  "word3",
			A:  "answer3",
			B: models.Backside{
				Type:  1,
				Value: "backside3",
			},
			DeckId: 1,
		},
		{
			Id: 4,
			W:  "word4",
			A:  "answer4",
			B: models.Backside{
				Type:  1,
				Value: "backside4",
			},
			DeckId: 1,
		},
		{
			Id: 5,
			W:  "word5",
			A:  "answer5",
			B: models.Backside{
				Type:  1,
				Value: "backside5",
			},
			DeckId: 1,
		},
	}
)

func (t *testingSuite) Test_GetNextRandomDeck() {
	t.m.On("GetUserInfo", mock.Anything, 1).
		Return(models.UserInfo{
			UserId: 1,
			MaxBox: 3,
		}, nil)
	t.m.On("GetFlashcardsIdByDeckId", mock.Anything, 1).
		Return(flashcard, nil)
	t.m.On("UpdateLeitner", mock.Anything, mock.Anything).Return(nil)
	t.m.On("GetLeitnerByUserIdCardId", mock.Anything, 1, mock.AnythingOfType("int")).
		Return(func(_ context.Context, i, j int) (models.UserLeitner, error) {
			if flashcardCalled[j-1] {
				return models.UserLeitner{
					Id:          0,
					UserId:      0,
					FlashcardId: 0,
					Box:         0,
					CoolDown:    models.CoolDown(time.Time{}.Add(time.Hour)),
				}, nil
			}
			return leitners[j-1], nil
		})

	cards := make([]models.FlashcardId, 0, 5)
	for i := 0; i < len(flashcard); i++ {
		t.Run(strconv.Itoa(i), func() {
			id, err := t.s.GetNextRandomDeck(context.Background(), 1, 1, models.CoolDown(time.Time{}.Add(time.Second)))
			if assert.NoError(t.T(), err) {
				cards = append(cards, id)
				flashcardCalled[id-1] = true
			}
		})
	}
	_, err := t.s.GetNextRandomDeck(context.Background(), 1, 1, models.CoolDown(time.Time{}.Add(time.Second)))
	assert.Error(t.T(), err)

	slices.Sort(cards)
	assert.Equal(t.T(), flashcard, cards)

	_, err = t.s.GetNextRandomDeck(context.Background(), 1, 1, models.CoolDown(time.Time{}.Add(2*time.Hour)))
	assert.NoError(t.T(), err)
}

func (t *testingSuite) Test_GetNextRandomDeckN() {
	flashcardCalled = make([]bool, 5)
	t.m.On("GetUserInfo", mock.Anything, 1).
		Return(models.UserInfo{
			UserId: 1,
			MaxBox: 3,
		}, nil)
	t.m.On("GetFlashcardsIdByDeckId", mock.Anything, 1).
		Return(flashcard, nil)
	t.m.On("UpdateLeitner", mock.Anything, mock.Anything).Return(nil)
	t.m.On("GetLeitnerByUserIdCardId", mock.Anything, 1, mock.AnythingOfType("int")).
		Return(func(_ context.Context, i, j int) (models.UserLeitner, error) {
			if flashcardCalled[j-1] {
				return models.UserLeitner{
					Id:          0,
					UserId:      0,
					FlashcardId: 0,
					Box:         0,
					CoolDown:    models.CoolDown(time.Time{}.Add(time.Hour)),
				}, nil
			}
			return leitners[j-1], nil
		})

	cards, err := t.s.GetNextRandomDeckN(context.Background(), 1, 1, models.CoolDown(time.Time{}.Add(time.Hour)), 5)
	assert.NoError(t.T(), err)
	assert.True(t.T(), len(cards) == 5)

	cards, err = t.s.GetNextRandomDeckN(context.Background(), 1, 1, models.CoolDown(time.Time{}.Add(time.Hour)), 3)
	assert.NoError(t.T(), err)
	assert.True(t.T(), len(cards) == 3)

	cards, err = t.s.GetNextRandomDeckN(context.Background(), 1, 1, models.CoolDown(time.Time{}.Add(time.Hour)), 100)
	assert.NoError(t.T(), err)
	assert.True(t.T(), len(cards) == 5)
}

func (t *testingSuite) Test_MakeMatchingDeck() {
	flashcardCalled = make([]bool, 5)
	t.m.On("GetUserInfo", mock.Anything, 1).
		Return(models.UserInfo{
			UserId: 1,
			MaxBox: 3,
		}, nil)
	t.m.On("GetFlashcardsIdByDeckId", mock.Anything, 1).
		Return(flashcard, nil)
	t.m.On("UpdateLeitner", mock.Anything, mock.Anything).Return(nil)
	t.m.On("GetLeitnerByUserIdCardId", mock.Anything, 1, mock.AnythingOfType("int")).
		Return(func(_ context.Context, i, j int) (models.UserLeitner, error) {
			if flashcardCalled[j-1] {
				return models.UserLeitner{
					Id:          0,
					UserId:      0,
					FlashcardId: 0,
					Box:         0,
					CoolDown:    models.CoolDown(time.Time{}.Add(time.Hour)),
				}, nil
			}
			return leitners[j-1], nil
		})
	t.m.On("GetNextRandomDeckN", mock.Anything, 1, 1, mock.Anything, 5).Return(flashcard)
	t.m.On("GetFlashcardById", mock.Anything, 1).Return(flashcardContents[0], nil)
	t.m.On("GetFlashcardById", mock.Anything, 2).Return(flashcardContents[1], nil)
	t.m.On("GetFlashcardById", mock.Anything, 3).Return(flashcardContents[2], nil)
	t.m.On("GetFlashcardById", mock.Anything, 4).Return(flashcardContents[3], nil)
	t.m.On("GetFlashcardById", mock.Anything, 5).Return(flashcardContents[4], nil)

	matching, err := t.s.MakeMatchingDeck(context.Background(), 1, 1, models.CoolDown(time.Time{}.Add(time.Hour)), 5)
	assert.NoError(t.T(), err)
	assert.True(t.T(), len(matching.Cards) == 5)
	answers := make(map[string]bool)
	for _, card := range matching.Cards {
		answers[card.A] = true
	}
	assert.True(t.T(), len(answers) == 5)
	for key, answer := range matching.M {
		_, ok := matching.Cards[key]
		assert.True(t.T(), ok)
		assert.True(t.T(), answers[answer])
	}
}

// func (t *testingSuite) Test_MakeTextDeck() {
// 	flashcardCalled = make([]bool, 5)
// 	t.m.On("GetUserInfo", mock.Anything, 1).
// 		Return(models.UserInfo{
// 			UserId: 1,
// 			MaxBox: 3,
// 		}, nil)
// 	t.m.On("GetFlashcardsIdByDeckId", mock.Anything, 1).
// 		Return(flashcard, nil)
// 	t.m.On("UpdateLeitner", mock.Anything, mock.Anything).Return(nil)
// 	t.m.On("GetLeitnerByUserIdCardId", mock.Anything, 1, mock.AnythingOfType("int")).
// 		Return(func(_ context.Context, i, j int) (models.UserLeitner, error) {
// 			if flashcardCalled[j-1] {
// 				return models.UserLeitner{
// 					Id:          0,
// 					UserId:      0,
// 					FlashcardId: 0,
// 					Box:         0,
// 					CoolDown:    models.CoolDown(time.Time{}.Add(time.Hour)),
// 				}, nil
// 			}
// 			return leitners[j-1], nil
// 		})
// 	t.m.On("GetNextRandomDeckN", mock.Anything, 1, 1, mock.Anything, 5).Return(flashcard)
// 	t.m.On("GetFlashcardById", mock.Anything, 1).Return(flashcardContents[0], nil)
// 	t.m.On("GetFlashcardById", mock.Anything, 2).Return(flashcardContents[1], nil)
// 	t.m.On("GetFlashcardById", mock.Anything, 3).Return(flashcardContents[2], nil)
// 	t.m.On("GetFlashcardById", mock.Anything, 4).Return(flashcardContents[3], nil)
// 	t.m.On("GetFlashcardById", mock.Anything, 5).Return(flashcardContents[4], nil)
// 	t.m.On("GenerateText", mock.Anything).Return(" asd sad word1 asdsa word2 word3 word4 word5 asdgsah sadhgashd ashjdgsad", nil)

// 	text, err := t.s.MakeTextDeck(context.Background(), 1, 1, models.CoolDown(time.Time{}.Add(time.Hour)), 5)
// 	assert.NoError(t.T(), err)
// 	assert.Equal(t.T(), ` asd sad <div style=\"display:none;\">word1</div> asdsa `+
// 		`<div style=\"display:none;\">word2</div> `+
// 		`<div style=\"display:none;\">word3</div> `+
// 		`<div style=\"display:none;\">word4</div> `+
// 		`<div style=\"display:none;\">word5</div> asdgsah sadhgashd ashjdgsad`, text)
// }

var leitners2 = []models.UserLeitner{
	{
		Id:          1,
		UserId:      1,
		FlashcardId: 1,
		Box:         0,
		CoolDown:    models.CoolDown{},
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

var flashcardContents2 = []models.Flashcard{
	{
		Id: 1,
		W:  "word1",
		A:  "answer1",
		B: models.Backside{
			Type:  1,
			Value: "backside1",
		},
		DeckId: 1,
	},
	{
		Id: 2,
		W:  "word2",
		A:  "answer2",
		B: models.Backside{
			Type:  1,
			Value: "backside2",
		},
		DeckId: 1,
	},
	{
		Id: 3,
		W:  "word3",
		A:  "answer3",
		B: models.Backside{
			Type:  1,
			Value: "backside3",
		},
		DeckId: 1,
	},
	{
		Id: 4,
		W:  "word4",
		A:  "answer4",
		B: models.Backside{
			Type:  1,
			Value: "backside4",
		},
		DeckId: 1,
	},
	{
		Id: 5,
		W:  "word5",
		A:  "answer5",
		B: models.Backside{
			Type:  1,
			Value: "backside5",
		},
		DeckId: 1,
	},
	{
		Id: 6,
		W:  "word6",
		A:  "answer6",
		B: models.Backside{
			Type:  1,
			Value: "backside6",
		},
		DeckId: 2,
	},
	{
		Id: 7,
		W:  "word7",
		A:  "answer7",
		B: models.Backside{
			Type:  1,
			Value: "backside7",
		},
		DeckId: 2,
	},
	{
		Id: 8,
		W:  "word8",
		A:  "answer8",
		B: models.Backside{
			Type:  1,
			Value: "backside8",
		},
		DeckId: 2,
	},
}

var (
	flashcard2       = []models.FlashcardId{6, 7, 8}
	flashcard2Called = [8]bool{}
)

const length = 8

func (t *testingSuite) Test_GetNextRandom() {
	t.m.On("GetDecksByUserId", mock.Anything, 1).Return(models.Decks{1: {}, 2: {}}, nil)
	t.m.On("GetUserInfo", mock.Anything, 1).
		Return(models.UserInfo{
			UserId: 1,
			MaxBox: 3,
		}, nil)
	t.m.On("GetFlashcardsIdByDeckId", mock.Anything, 1).
		Return(flashcard, nil)
	t.m.On("GetFlashcardsIdByDeckId", mock.Anything, 2).
		Return(flashcard2, nil)
	t.m.On("UpdateLeitner", mock.Anything, mock.Anything).Return(nil)
	t.m.On("GetLeitnerByUserIdCardId", mock.Anything, 1, mock.AnythingOfType("int")).
		Return(func(_ context.Context, i, j int) (models.UserLeitner, error) {
			if flashcard2Called[j-1] {
				return models.UserLeitner{
					Id:          0,
					UserId:      0,
					FlashcardId: 0,
					Box:         0,
					CoolDown:    models.CoolDown(time.Time{}.Add(time.Hour)),
				}, nil
			}
			return leitners2[j-1], nil
		})

	cards := make([]models.FlashcardId, 0, length)
	for i := 0; i < length; i++ {
		t.Run(strconv.Itoa(i), func() {
			id, err := t.s.GetNextRandom(context.Background(), 1, models.CoolDown(time.Time{}.Add(time.Second)))
			if assert.NoError(t.T(), err) {
				cards = append(cards, id)
				flashcard2Called[id-1] = true
			}
		})
	}
	_, err := t.s.GetNextRandom(context.Background(), 1, models.CoolDown(time.Time{}.Add(time.Second)))
	assert.Error(t.T(), err)

	slices.Sort(cards)
	assert.Equal(t.T(), append(flashcard, flashcard2...), cards)

	_, err = t.s.GetNextRandom(context.Background(), 1, models.CoolDown(time.Time{}.Add(2*time.Hour)))
	assert.NoError(t.T(), err)
}

func (t *testingSuite) Test_GetNextRandomN() {
	flashcard2Called = [8]bool{}
	t.m.On("GetDecksByUserId", mock.Anything, 1).Return(models.Decks{1: {}, 2: {}}, nil)
	t.m.On("GetUserInfo", mock.Anything, 1).
		Return(models.UserInfo{
			UserId: 1,
			MaxBox: 3,
		}, nil)
	t.m.On("GetFlashcardsIdByDeckId", mock.Anything, 1).
		Return(flashcard, nil)
	t.m.On("GetFlashcardsIdByDeckId", mock.Anything, 2).
		Return(flashcard2, nil)
	t.m.On("UpdateLeitner", mock.Anything, mock.Anything).Return(nil)
	t.m.On("GetLeitnerByUserIdCardId", mock.Anything, 1, mock.AnythingOfType("int")).
		Return(func(_ context.Context, i, j int) (models.UserLeitner, error) {
			if flashcard2Called[j-1] {
				return models.UserLeitner{
					Id:          0,
					UserId:      0,
					FlashcardId: 0,
					Box:         0,
					CoolDown:    models.CoolDown(time.Time{}.Add(time.Hour)),
				}, nil
			}
			return leitners2[j-1], nil
		})

	cards, err := t.s.GetNextRandomN(context.Background(), 1, models.CoolDown(time.Time{}.Add(time.Second)), 8)
	assert.NoError(t.T(), err)
	assert.True(t.T(), len(cards) == 8)

	cards, err = t.s.GetNextRandomN(context.Background(), 1, models.CoolDown(time.Time{}.Add(time.Second)), 3)
	assert.NoError(t.T(), err)
	assert.True(t.T(), len(cards) == 3)

	cards, err = t.s.GetNextRandomN(context.Background(), 1, models.CoolDown(time.Time{}.Add(time.Second)), 100)
	assert.NoError(t.T(), err)
	assert.True(t.T(), len(cards) == 8)
}

func (t *testingSuite) Test_MakeMatching() {
	flashcard2Called = [8]bool{}

	t.m.On("GetDecksByUserId", mock.Anything, 1).Return(models.Decks{1: {}, 2: {}}, nil)
	t.m.On("GetUserInfo", mock.Anything, 1).
		Return(models.UserInfo{
			UserId: 1,
			MaxBox: 3,
		}, nil)
	t.m.On("GetFlashcardsIdByDeckId", mock.Anything, 1).
		Return(flashcard, nil)
	t.m.On("GetFlashcardsIdByDeckId", mock.Anything, 2).
		Return(flashcard2, nil)
	t.m.On("UpdateLeitner", mock.Anything, mock.Anything).Return(nil)
	t.m.On("GetLeitnerByUserIdCardId", mock.Anything, 1, mock.AnythingOfType("int")).
		Return(func(_ context.Context, i, j int) (models.UserLeitner, error) {
			if flashcard2Called[j-1] {
				return models.UserLeitner{
					Id:          0,
					UserId:      0,
					FlashcardId: 0,
					Box:         0,
					CoolDown:    models.CoolDown(time.Time{}.Add(time.Hour)),
				}, nil
			}
			return leitners2[j-1], nil
		})
	var flashcardIds = []models.FlashcardId{1, 2, 3, 4, 5, 6, 7, 8}
	t.m.On("GetNextRandomN", mock.Anything, 1, 1, mock.Anything, 8).Return(flashcardIds)
	t.m.On("GetFlashcardById", mock.Anything, 1).Return(flashcardContents2[0], nil)
	t.m.On("GetFlashcardById", mock.Anything, 2).Return(flashcardContents2[1], nil)
	t.m.On("GetFlashcardById", mock.Anything, 3).Return(flashcardContents2[2], nil)
	t.m.On("GetFlashcardById", mock.Anything, 4).Return(flashcardContents2[3], nil)
	t.m.On("GetFlashcardById", mock.Anything, 5).Return(flashcardContents2[4], nil)
	t.m.On("GetFlashcardById", mock.Anything, 6).Return(flashcardContents2[5], nil)
	t.m.On("GetFlashcardById", mock.Anything, 7).Return(flashcardContents2[6], nil)
	t.m.On("GetFlashcardById", mock.Anything, 8).Return(flashcardContents2[7], nil)

	matching, err := t.s.MakeMatching(context.Background(), 1, models.CoolDown(time.Time{}.Add(time.Hour)), 8)
	assert.NoError(t.T(), err)
	assert.True(t.T(), len(matching.Cards) == 8)
	answers := make(map[string]bool)
	for _, card := range matching.Cards {
		answers[card.A] = true
	}
	assert.True(t.T(), len(answers) == 8)
	for key, answer := range matching.M {
		_, ok := matching.Cards[key]
		assert.True(t.T(), ok)
		assert.True(t.T(), answers[answer])
	}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(testingSuite))
}
