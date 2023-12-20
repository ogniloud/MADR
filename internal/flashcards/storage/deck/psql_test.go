//go:build psql_test
// +build psql_test

package deck

import (
	"context"
	"io"
	"sort"
	"testing"
	"time"

	"github.com/charmbracelet/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/ogniloud/madr/internal/db"
	"github.com/ogniloud/madr/internal/flashcards/models"
)

type PSQLSuite struct {
	suite.Suite
	repo *Storage
}

const connStringTest = "user=postgres password=postgres host=localhost port=5432 dbname=postgres"

func (s *PSQLSuite) SetupSuite() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := db.NewPSQLDatabase(ctx, connStringTest, log.New(io.Discard))
	if err != nil {
		s.T().Fatal(err)
	}

	s.repo = &Storage{Conn: conn}
}

func (s *PSQLSuite) TearDownSuite() {
	s.repo.Conn.Close()
}

func (s *PSQLSuite) TestDeckStorage_GetDecksByUserId() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	decks, err := s.repo.GetDecksByUserId(ctx, 1488)

	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), 2, len(decks))
		keys := decks.Keys()
		sort.Ints(keys)
		assert.Equal(s.T(), keys, []models.DeckId{1, 2})
	}
}

func (s *PSQLSuite) TestDeckStorage_GetFlashcardsIdByDeckId() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cards, err := s.repo.GetFlashcardsIdByDeckId(ctx, 1)

	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), 2, len(cards))
		assert.Equal(s.T(), cards, []models.FlashcardId{1, 2})
	}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(PSQLSuite))
}
