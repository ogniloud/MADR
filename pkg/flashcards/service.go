package flashcards

import (
	"fmt"
	"github.com/ogniloud/madr/pkg/flashcards/leitner"
	"github.com/ogniloud/madr/pkg/flashcards/storage"
	"io"
	"runtime"
)

type Service interface {
	SetLeitner(userId int) error
	Release() error

	GetNextFlashcard() (*leitner.Flashcard, error)
	GetNextFlashcardDeck(id leitner.DeckId) (*leitner.Flashcard, error)
	RateFlashcard(id leitner.CardId, rate leitner.Rate) error

	// todo

	MakeText() (io.Reader, error)
	MakeTextDeck(id leitner.DeckId) (io.Reader, error)

	MakeExercise()
	MakeExerciseDeck(id leitner.DeckId)

	PublishDeck(id leitner.DeckId) error
}

type LeitnerService struct {
	cs  storage.CardStorer
	dps storage.DeckPrivacyStorer
	ls  storage.LeitnerStorer
	ucs storage.UserConfigStorer

	l *leitner.Leitner
}

func (ls LeitnerService) SetLeitner(userId int) error {
	ltns, err := ls.ls.GetAllByUserId(userId)
	if err != nil {
		return fmt.Errorf("failed to load cards for userid %v: %w", userId, err)
	}

	cfg, err := ls.ucs.Get(userId)
	if err != nil {
		return fmt.Errorf("failed to load user config %v: %w", userId, err)
	}

	// collect decks for loading them to Leitner
	// map is needed for creating new maps with different id's (id can be very big)
	// decks will be passed to NewLeitner func
	// so, invariant is len(m) == len(decks)
	decks := make([]leitner.Deck, 0)
	m := make(map[leitner.DeckId]*leitner.Deck)

	for _, ltn := range ltns {
		// load a simple card
		cardDB, err := ls.cs.Get(ltn.CardId)
		if err != nil {
			return fmt.Errorf("failed to load a flashcard for userid %v: %w", userId, err)
		}

		deck, ok := m[cardDB.DeckId]
		if !ok {
			decktmp := leitner.NewDeck(cfg.MaxLevel)
			decks = append(decks, decktmp)

			deck = &decktmp
			m[cardDB.DeckId] = deck
		}

		err = deck.Insert(&leitner.Flashcard{
			Id: cardDB.Id,
			W:  cardDB.Word,
			B:  cardDB.Backside,
			L:  ltn.Level,
			Cd: ltn.Cd,
		})
		if err != nil {
			return fmt.Errorf("failed to insert a flashcard to deck for userid %v: %w", userId, err)
		}
	}

	l := leitner.NewLeitner(cfg.MaxLevel, decks, func(level leitner.Level) leitner.CoolDown {
		return leitner.CoolDownFunc(level)
	})
	ls.l = &l

	return nil
}

func (ls LeitnerService) Release() error {
	if ls.l == nil {
		return fmt.Errorf("already released")
	}

	ls.l = nil
	runtime.GC()

	return nil
}

func (ls LeitnerService) GetNextFlashcard() (*leitner.Flashcard, error) {
	if ls.l == nil {
		return nil, fmt.Errorf("nil leitner")
	}

	fc, err := ls.l.GetRandom()
	if err != nil {
		return nil, fmt.Errorf("failed to get next card: %w", err)
	}

	return fc, nil
}

func (ls LeitnerService) GetNextFlashcardDeck(id leitner.DeckId) (*leitner.Flashcard, error) {
	if ls.l == nil {
		return nil, fmt.Errorf("nil leitner")
	}

	deck, err := ls.l.Deck(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get a deck for next falshcard: %w", err)
	}

	return deck.GetRandom(ls.l.Probabilities())
}

func (ls LeitnerService) RateFlashcard(id leitner.CardId, rate leitner.Rate) error {
	if ls.l == nil {
		return fmt.Errorf("nil leitner")
	}

	fcDB, err := ls.cs.Get(id)
	if err != nil {
		return fmt.Errorf("failed to get a flashcard and rate: %w", err)
	}

	deck, _ := ls.l.Deck(fcDB.DeckId)

	for _, box := range deck.Boxes {
		if fc, err := box.Get(id); err == nil {
			return ls.l.Rate(fc, rate)
		}
	}

	return fmt.Errorf("failed to rate a flashcard: %w", err)
}
