package psql

import (
	"context"
	"io"
	"testing"

	"github.com/charmbracelet/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/ogniloud/madr/internal/db"
	"github.com/ogniloud/madr/internal/flashcards/models"
)

type PSQLSuite struct {
	suite.Suite
	repo *DeckStorage
}

const connStringTest = "user=postgres password=postgres host=localhost port=5432 dbname=postgres"

func (s *PSQLSuite) SetupSuite() {
	conn, err := db.NewPSQLDatabase(context.Background(), connStringTest, log.New(io.Discard))
	if err != nil {
		s.T().Fatal(err)
	}

	s.repo = &DeckStorage{Conn: conn}
}

func (s *PSQLSuite) TearDownSuite() {
	_ = s.repo.Conn.Close(context.Background())
}

func (s *PSQLSuite) TestDeckStorage_GetDecksByUserId() {
	decks, err := s.repo.GetDecksByUserId(1)

	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), 2, len(decks))
		assert.Equal(s.T(), decks.Keys(), []models.DeckId{12, 23})
	}
}

func (s *PSQLSuite) TestDeckStorage_GetFlashcardsIdByDeckId() {
	cards, err := s.repo.GetFlashcardsIdByDeckId(1)

	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), 2, len(cards))
		assert.Equal(s.T(), cards, []models.FlashcardId{1, 2})
	}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(PSQLSuite))
}
