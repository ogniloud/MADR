package study

import (
	"fmt"
	"math"
	"math/rand"
	"slices"

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

// GetNextRandom возвращает случаёную карточку из всего набора карточек пользователя с истёкшим CoolDown.
func (s *StudyService) GetNextRandom(uid models.UserId, down models.CoolDown) (models.FlashcardId, error) {
	decks, err := s.fserv.LoadDecks(uid)
	if err != nil {
		return 0, err
	}

	keys := decks.Keys()
	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	for i := 0; i < len(keys); i++ {
		card, err := s.GetNextRandomDeck(uid, keys[i], down)
		if err == nil {
			return card, nil
		}
	}

	return 0, fmt.Errorf("no cards")
}

// GetNextRandomDeck возвращает случайную карточку из колоды пользователя с истёкшим CoolDown.
func (s *StudyService) GetNextRandomDeck(uid models.UserId, id models.DeckId, down models.CoolDown) (models.FlashcardId, error) {
	ids, err := s.fserv.GetFlashcardsIdByDeckId(id)
	if err != nil {
		return 0, err
	}
	ltns := make([]models.UserLeitner, 0, len(ids))

	for _, id := range ids {
		ltn, err := s.fserv.GetLeitnerByUserIdCardId(uid, id)
		if err != nil {
			return 0, err
		}

		ltns = append(ltns, ltn)
	}

	ltns = slices.DeleteFunc(ltns, func(leitner models.UserLeitner) bool {
		return !leitner.CoolDown.IsPassed(down)
	})

	if len(ltns) == 0 {
		return 0, fmt.Errorf("no cards")
	}

	// shuffle and choose next card
	rand.Shuffle(len(ltns), func(i, j int) {
		ltns[i], ltns[j] = ltns[j], ltns[i]
	})
	b := s.box()

	for _, v := range ltns {
		if v.Box == b {
			return v.FlashcardId, nil
		}
	}
	return ltns[rand.Intn(len(ltns))].FlashcardId, nil
}

// Rate перемещает карточку в бокс относительно оценки.
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
