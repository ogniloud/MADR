package leitner

import (
	"github.com/ogniloud/madr/pkg/flashcards/types"
	"testing"
)

type CoolDownTest int

var t = CoolDownTest(0)

func (cd CoolDownTest) Passed() bool {
	return t >= cd
}

func cooldownTest(l types.Level) types.CoolDown {
	switch l {
	case 0:
		return t + 1
	case 1:
		return t + 2
	case 2:
		return t + 5
	}
	panic("invalid level")
}

func TestCoolDownPassed(t *testing.T) {
	cards := []*types.Flashcard{
		{
			Id: 0,
			W:  "Zero",
			B:  Translation("Ноль"),
			L:  0,
			Cd: CoolDownTest(0),
		}, {
			Id: 1,
			W:  "One",
			B:  Translation("Адин"),
			L:  0,
			Cd: CoolDownTest(0),
		}, {
			Id: 2,
			W:  "Two",
			B:  Translation("Два"),
			L:  1,
			Cd: CoolDownTest(0),
		}, {
			Id: 3,
			W:  "Three",
			B:  Translation("Три"),
			L:  2,
			Cd: CoolDownTest(0),
		},
	}

	b1 := types.Box(NewBox(0))
	b2 := types.Box(NewBox(1))
	b3 := types.Box(NewBox(2))

	d1 := types.Deck(Deck{
		maxLevel: 3,
		boxes:    []types.Box{b1, b2, b3},
	})
	d2 := types.Deck(Deck{
		maxLevel: 3,
		boxes:    []types.Box{b1, b2, b3},
	})

	if err := d1.Insert(cards[0]); err != nil {
		t.Errorf("unexpected error %v", err)
		return
	}
	if err := d1.Insert(cards[1]); err != nil {
		t.Errorf("unexpected error %v", err)
		return
	}
	if err := d2.Insert(cards[2]); err != nil {
		t.Errorf("unexpected error %v", err)
		return
	}
	if err := d1.Insert(cards[3]); err != nil {
		t.Errorf("unexpected error %v", err)
		return
	}
	lt := NewLeitner(3, []types.Deck{d1, d2}, cooldownTest)

	if _, _, err := lt.GetRandom(); err != nil {
		t.Errorf("unexpected error %v", err)
		return
	}

	for _, c := range cards[:3] {
		c.CoolDown(cooldownTest)
	}

	fc, _, err := lt.GetRandom()
	if err != nil {
		t.Errorf("unexpected error %v", err)
		return
	}

	if fc.Id != 3 {
		t.Errorf("Bad random: Expected id=3, got id=%v", fc.Id)
		return
	}
	cards[3].CoolDown(cooldownTest)

	fc, _, err = lt.GetRandom()
	if err == nil {
		t.Errorf("wanted error but got nil, fc: %#+v", *fc)
		return
	}
}
