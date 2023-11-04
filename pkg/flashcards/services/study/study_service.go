package study

import (
	"fmt"
	"math"
	"math/rand"
	"slices"

	"github.com/ogniloud/madr/pkg/flashcards/models"
	"github.com/ogniloud/madr/pkg/flashcards/services/deck"
)

var ErrNoCards = fmt.Errorf("no cards")

// Mark is a type of rating marks of flashcards in Leitner's system.
type Mark int

const (
	Bad = Mark(iota)
	Satisfactory
	Excellent
)

type IStudyService interface {
	// GetNextRandom returns a random card from the entire set of user cards with expired CoolDown.
	GetNextRandom(uid models.UserId, down models.CoolDown) (models.FlashcardId, error)

	// GetNextRandomDeck returns a random card from the user's deck whose CoolDown has expired.
	GetNextRandomDeck(uid models.UserId, id models.DeckId, down models.CoolDown) (models.FlashcardId, error)

	// Rate moves the card into the box relative to the mark.
	Rate(uid models.UserId, id models.FlashcardId, mark Mark) error
}

type Service struct {
	// dsrv stands for deck.Service
	dsrv deck.IService

	// p is a probability distribution for card selection functions,
	// the list must be equal to the maximum number of boxes the user has.
	p []float32 // for each i: 0 <= p[i] < 1
	// Example: Let p = [0, 0,5, 0.9] - three boxes: 0, 1, 2.
	// |____________<_______<__)
	// 0           0.5     0.9 1
	//
	// r = 0.6
	// Since p[1] <= r < p[2], we choose the flashcard from the box 1.
}

func NewService(s deck.IService, maxBox models.Box) IStudyService {
	p := make([]float32, maxBox)
	p[0] = 0

	for i := range p[1:] {
		p[i+1] = p[i] + float32(math.Pow(.5, float64(i+1)))
	}

	return &Service{dsrv: s, p: p}
}

// GetNextRandom returns a random card from the entire set of user cards with expired CoolDown.
func (s *Service) GetNextRandom(uid models.UserId, down models.CoolDown) (models.FlashcardId, error) {
	decks, err := s.dsrv.LoadDecks(uid)
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

	return 0, ErrNoCards
}

// GetNextRandomDeck returns a random card from the user's deck whose CoolDown has expired.
func (s *Service) GetNextRandomDeck(uid models.UserId, id models.DeckId, down models.CoolDown) (models.FlashcardId, error) {
	ids, err := s.dsrv.GetFlashcardsIdByDeckId(id)
	if err != nil {
		return 0, err
	}

	ltns := make([]models.UserLeitner, 0, len(ids))

	for _, id := range ids {
		ltn, err := s.dsrv.GetLeitnerByUserIdCardId(uid, id)
		if err != nil {
			return 0, err
		}

		ltns = append(ltns, ltn)
	}

	ltns = slices.DeleteFunc(ltns, func(leitner models.UserLeitner) bool {
		return !leitner.CoolDown.IsPassed(down)
	})

	if len(ltns) == 0 {
		return 0, ErrNoCards
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

// Rate moves the card into the box relative to the mark.
func (s *Service) Rate(uid models.UserId, id models.FlashcardId, mark Mark) error {
	l, err := s.dsrv.GetLeitnerByUserIdCardId(uid, id)
	if err != nil {
		return err
	}

	box, err := s.dsrv.UserMaxBox(uid)
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

	return s.dsrv.UpdateLeitner(l)
}

// let p = [.5 .7 .8 .9 .95] and r is a random float
// returned Box is the minimum index i such that r >= p[i]
// so, for each i => 0 <= p[i] < 1.
func (s *Service) box() models.Box {
	r := rand.Float32()
	for i, v := range s.p {
		if r >= v && (i == len(s.p)-1 || s.p[i+1] > r) {
			return i
		}
	}

	return len(s.p)
}
