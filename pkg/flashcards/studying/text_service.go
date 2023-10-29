package studying

import (
	"io"

	"github.com/ogniloud/madr/pkg/flashcards/storage"
)

type ITextService interface {
	GenerateText(flashcards []storage.FlashcardId) (io.Reader, error)
	GenerateTextDeck(flashcards []storage.FlashcardId, id storage.DeckId) (io.Reader, error)
}
