package studying

import (
	"github.com/ogniloud/madr/pkg/flashcards/storage"
	"io"
)

type Mark int

const (
	Bad = Mark(iota)
	Satisfactory
	Excellent
)

type TextService interface {
	GenerateText(flashcards []storage.FlashcardId) (io.Reader, error)
	GenerateTextDeck(flashcards []storage.FlashcardId, id storage.DeckId) (io.Reader, error)
}

type StudyService interface {
	GetNRandom(n int) ([]storage.FlashcardId, error) // n > 0
	GetNRandomDeck(n int, id storage.DeckId) ([]storage.FlashcardId, error)
	Rate(id storage.FlashcardId, mark Mark) error
}

type Matching struct {
}

type Sentence struct {
}

type ExerciseService interface {
	MakeMatching() Matching
	MakeSentence() Sentence
}
