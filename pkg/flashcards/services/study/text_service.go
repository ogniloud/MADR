package study

import (
	"io"

	"github.com/ogniloud/madr/pkg/flashcards/models"
)

type ITextService interface {
	GenerateText(flashcards []models.FlashcardId) (io.Reader, error)
	GenerateTextDeck(flashcards []models.FlashcardId, id models.DeckId) (io.Reader, error)
}
