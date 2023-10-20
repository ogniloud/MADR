package leitner

import (
	"math"
	"math/rand"
)

type Id int

type Rate int

const (
	Bad = Rate(iota)
	Satisfactory
	Good
)

// Leitner is an abstract data structure consisting of
// Boxes. Each box has a temperature Level. It means that
// the hotter the box the higher chance to be chosen by GetRandom.
// Leitner should consist of deck (deck of decks).
//
// Also, for each flashcard defined CoolDown. A flashcard can't be returned
// if CoolDown is not passed.
type Leitner struct {
	maxLevel Level // 0 to maxLevel-1
	decks    []Deck
	cooldown func(Level) CoolDown
	p        []float32
}

func NewLeitner(
	maxLevel Level,
	decks []Deck,
	cooldown func(Level) CoolDown,
) Leitner {
	p := countDistribution(0.2, maxLevel)
	l := Leitner{p: p, maxLevel: maxLevel, decks: decks, cooldown: cooldown}

	return l
}

// Rate takes a mark from the user and inserts the card
// in a corresponding to its level box
func (l Leitner) Rate(fc *Flashcard, rate Rate) error {
	if rate == Satisfactory {
		fc.CoolDown(l.cooldown)
	}

	d, err := l.Deck(fc.DeckId)
	if err != nil {
		return err
	}

	err = d.Delete(fc.Id)
	if err != nil {
		return err
	}

	if rate == Good && fc.L < l.maxLevel-1 {
		fc.L++
	} else if rate == Bad && fc.L > 0 {
		fc.L--
	}

	fc.CoolDown(l.cooldown)

	return d.Insert(fc)
}

// GetRandom means take a random card from a random deck
func (l Leitner) GetRandom() (*Flashcard, error) {
	ri := DeckId(rand.Intn(len(l.decks)))
	deck, _ := l.Deck(ri)

	fc, err := deck.GetRandom(l.p)
	if err == nil {
		return fc, nil
	}

	ln := DeckId(len(l.decks))
	for j := (ri + 1) % ln; j != ri; j = (j + 1) % ln {
		deck, _ = l.Deck(j)
		fc, err = deck.GetRandom(l.p)
		if err == nil {
			return fc, nil
		}
	}

	return nil, ErrCardsUnavailable
}

func (l Leitner) Deck(id DeckId) (Deck, error) {
	if id < 0 || int(id) >= len(l.decks) {
		return Deck{}, ErrDeckBadIndex
	}

	return l.decks[id], nil
}

func (l Leitner) Probabilities() []float32 {
	return l.p
}

func countDistribution(offset float32, maxLevel Level) []float32 {
	p := make([]float32, maxLevel)
	rest := 1 - offset

	for i := 1; i <= int(maxLevel); i++ {
		p[i-1] = offset + rest*float32(math.Pow(0.5, float64(i)))
		offset = p[i-1]
	}

	return p
}
