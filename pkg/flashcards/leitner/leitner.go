package leitner

import (
	ftypes "github.com/ogniloud/madr/pkg/flashcards/types"
	"math"
	"math/rand"
)

type Leitner struct {
	maxLevel ftypes.Level // 0 to maxLevel-1
	decks    []ftypes.Deck
	cooldown func(ftypes.Level) ftypes.CoolDown
	p        []float32
}

func (l Leitner) Rate(fc *ftypes.Flashcard, id ftypes.DeckId, rate ftypes.Rate) error {
	if rate == ftypes.Satisfactory {
		fc.CoolDown(l.cooldown)
	}

	d, err := l.Deck(id)
	if err != nil {
		return err
	}

	err = d.Delete(fc.Id)
	if err != nil {
		return err
	}

	if rate == ftypes.Good && fc.L < l.maxLevel-1 {
		fc.L++
	} else if rate == ftypes.Bad && fc.L > 0 {
		fc.L--
	}

	fc.CoolDown(l.cooldown)

	return d.Insert(fc)
}

func (l Leitner) GetRandom() (*ftypes.Flashcard, ftypes.DeckId, error) {
	ri := ftypes.DeckId(rand.Intn(len(l.decks)))
	deck, _ := l.Deck(ri)

	fc, err := deck.GetRandom(l.p)
	if err == nil {
		return fc, ri, nil
	}

	ln := ftypes.DeckId(len(l.decks))
	for j := (ri + 1) % ln; j != ri; j = (j + 1) % ln {
		deck, _ = l.Deck(j)
		fc, err = deck.GetRandom(l.p)
		if err == nil {
			return fc, ri, nil
		}
	}

	return nil, 0, ErrBoxBadIndex
}

func (l Leitner) Deck(id ftypes.DeckId) (ftypes.Deck, error) {
	if id < 0 || int(id) >= len(l.decks) {
		return nil, ErrDeckBadIndex
	}

	return l.decks[id], nil
}

func NewLeitner(
	maxLevel ftypes.Level,
	decks []ftypes.Deck,
	cooldown func(ftypes.Level) ftypes.CoolDown,
) Leitner {
	p := countDistribution(0.2, maxLevel)
	l := Leitner{p: p, maxLevel: maxLevel, decks: decks, cooldown: cooldown}

	return l
}

func countDistribution(offset float32, maxLevel ftypes.Level) []float32 {
	p := make([]float32, maxLevel)
	rest := 1 - offset

	for i := 1; i <= int(maxLevel); i++ {
		p[i-1] = offset + rest*float32(math.Pow(0.5, float64(i)))
		offset = p[i-1]
	}

	return p
}
