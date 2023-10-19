package flashcards

import (
	"github.com/ogniloud/madr/pkg/flashcards/leitner"
	"github.com/ogniloud/madr/pkg/flashcards/storage"
)

type Service interface {
	SetLeitner(storage.UserId) error
	GetNextFlashcard() (*leitner.Flashcard, error)
	GetNextFlashcardDeck(*leitner.Deck) (*leitner.Flashcard, error)
}

type LeitnerService struct {
	Ds storage.DeckStorer
	Ls storage.LeitnerStorer
	L  *leitner.Leitner
}

func (s LeitnerService) SetLeitnerFromDecks(userId storage.UserId) error {
	lts, err := s.Ls.GetAllByUserId(userId)
	if err != nil {
		return err
	}

	decks := make([]leitner.Deck, 0, len(lts))

	for k, lt := range lts {
		decks = append(decks, leitner.NewDeck(lt.Level))

		deck, err := s.Ds.Get(lt.DkId)
		if err != nil {
			return err
		}
		for i, card := range deck.Fcs {
			fc := &leitner.Flashcard{
				Id: leitner.CardId(i),
				W:  card.W,
				B:  card.B,
				L:  lt.Decks[i],
				Cd: card.Cd,
			}
			if err := decks[k].Insert(fc); err != nil {
				return err
			}
		}
	}

	newLt := leitner.NewLeitner(lts[0].Level, decks, func(l leitner.Level) leitner.CoolDown {
		return leitner.CoolDownFunc(l)
	})
	s.L = &newLt

	return nil
}

func (s LeitnerService) GetNextFlashcard() (*leitner.Flashcard, error) {
	//TODO implement me
	panic("implement me")
}

func (s LeitnerService) GetNextFlashcardDeck(deck *leitner.Deck) (*leitner.Flashcard, error) {
	//TODO implement me
	panic("implement me")
}

type Handler struct {
	s Service
}
