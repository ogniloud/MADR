package studying

import (
	"math/rand"
	"time"

	"github.com/ogniloud/madr/pkg/flashcards"
	"github.com/ogniloud/madr/pkg/flashcards/storage"
)

type Mark int

const (
	Bad = Mark(iota)
	Satisfactory
	Excellent
)

type IStudyService interface {
	GetNRandom(uid storage.UserId, n int) ([]storage.FlashcardId, error) // n > 0
	GetNRandomDeck(uid storage.UserId, n int, id storage.DeckId) ([]storage.FlashcardId, error)
	Rate(id storage.FlashcardId, mark Mark) error
}

type StudyService struct {
	fserv flashcards.Service
	p     []float32
}

func (s StudyService) GetNRandom(uid storage.UserId, n int) ([]storage.FlashcardId, error) {
	m := map[storage.Box]int{}

	for i := 0; i < n; i++ {
		m[s.box()]++
	}

	return s.fserv.GetRandom(uid, storage.CoolDown{State: time.Now()}, m)
}

func (s StudyService) GetNRandomDeck(uid storage.UserId, n int, id storage.DeckId) ([]storage.FlashcardId, error) {
	m := map[storage.Box]int{}

	for i := 0; i < n; i++ {
		m[s.box()]++
	}

	return s.fserv.GetRandomDeck(uid, storage.CoolDown{State: time.Now()}, id, m)
}

func (s StudyService) Rate(uid storage.UserId, id storage.FlashcardId, mark Mark) error {
	l, err := s.fserv.S.GetLeitnerByUserIdCardId(uid, id)
	if err != nil {
		return err
	}

	box, err := s.fserv.GetUserMaxBox(uid)
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

	return s.fserv.S.UpdateLeitner(l)
}

func (s StudyService) box() storage.Box {
	r := rand.Float32()
	for i, v := range s.p {
		if r >= v {
			return i
		}
	}
	return len(s.p)
}
