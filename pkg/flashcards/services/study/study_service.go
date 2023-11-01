package study

import (
	"math"
	"math/rand"

	"github.com/ogniloud/madr/pkg/flashcards/models"
	"github.com/ogniloud/madr/pkg/flashcards/services/deck"
)

type Mark int

const (
	Bad = Mark(iota)
	Satisfactory
	Excellent

	CacheRandomLoadSize = 50
)

type IStudyService interface {
	GetNRandom(uid models.UserId, n int) ([]models.FlashcardId, error) // n > 0
	GetNRandomDeck(uid models.UserId, n int, id models.DeckId) ([]models.FlashcardId, error)
	Rate(id models.FlashcardId, mark Mark) error
}

type StudyService struct {
	fserv *deck.Service
	p     []float32 // for each i => 0 <= p[i] < 1
}

func NewStudy(s *deck.Service, maxBox models.Box) *StudyService {
	p := make([]float32, maxBox)
	p[0] = 1
	for i := range p[1:] {
		p[i+1] = p[i] + float32(math.Pow(.5, float64(i+1)))
	}

	return &StudyService{fserv: s, p: p}
}

func (s *StudyService) GetNextRandom(uid models.UserId, n int) (models.FlashcardId, error) {
	return 0, nil
}

func (s *StudyService) GetNextRandomDeck(uid models.UserId, n int, id models.DeckId) (models.FlashcardId, error) {
	return 0, nil
}

func (s *StudyService) Rate(uid models.UserId, id models.FlashcardId, mark Mark) error {
	l, err := s.fserv.GetLeitnerByUserIdCardId(uid, id)
	if err != nil {
		return err
	}

	box, err := s.fserv.UserMaxBox(uid)
	if err != nil {
		return err
	}

	switch mark {
	case Bad:
		if l.Box != 0 {
			l.Box--
		}
	case Satisfactory:
	case Excellent:
		if l.Box != box {
			l.Box++
		}
	}

	return s.fserv.UpdateLeitner(l)
}

// let p = [.5 .7 .8 .9 .95] and r is a random float
// returned Box is the minimum index i such that r >= p[i]
// so, for each i => 0 <= p[i] < 1
func (s *StudyService) box() models.Box {
	r := rand.Float32()
	for i, v := range s.p {
		if r >= v && (i == len(s.p)-1 || s.p[i+1] > r) {
			return i
		}
	}
	return len(s.p)
}
