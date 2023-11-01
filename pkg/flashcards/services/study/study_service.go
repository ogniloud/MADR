package study

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/ogniloud/madr/pkg/flashcards/cache"
	"github.com/ogniloud/madr/pkg/flashcards/models"
	"github.com/ogniloud/madr/pkg/flashcards/services/deck"
)

type Mark int

const (
	Bad = Mark(iota)
	Satisfactory
	Excellent
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

func (s *StudyService) GetRandom(uid models.UserId, n int) (models.FlashcardId, error) {
	d, ok := s.fserv.Cache().Load(fmt.Sprintf("k%d", uid))
	decks := d.(cache.CachedRandom)

	if !ok || len(decks) == 0 {
		m := map[models.Box]int{}

		for i := 0; i < n; i++ {
			m[s.box()]++
		}

		err := s.fserv.Random(uid, models.CoolDown{State: time.Now()}, m)
		if err != nil {
			return -1, err
		}
	}
	defer func() {
		decks = decks[1:]
		s.fserv.Cache().Store(fmt.Sprintf("k%d", uid), decks)
	}()

	return decks[0], nil
}

func (s *StudyService) GetRandomDeck(uid models.UserId, n int, id models.DeckId) (models.FlashcardId, error) {
	d, ok := s.fserv.Cache().Load(fmt.Sprintf("%dk%d", id, uid))
	decks := d.(cache.CachedRandom)

	if !ok || len(decks) == 0 {
		m := map[models.Box]int{}

		for i := 0; i < n; i++ {
			m[s.box()]++
		}

		err := s.fserv.RandomDeck(uid, models.CoolDown{State: time.Now()}, id, m)
		if err != nil {
			return -1, err
		}
	}
	defer func() {
		decks = decks[1:]
		s.fserv.Cache().Store(fmt.Sprintf("%dk%d", id, uid), decks)
	}()

	return decks[0], nil
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
