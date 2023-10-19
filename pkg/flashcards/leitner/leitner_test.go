package leitner

import (
	"testing"
)

type CoolDownTest int

var tm = CoolDownTest(0)

func (cd CoolDownTest) Passed() bool {
	return tm >= cd
}

func cooldownTest(l Level) CoolDown {
	switch l {
	case 0:
		return tm + 1
	case 1:
		return tm + 2
	case 2:
		return tm + 5
	}
	panic("invalid level")
}

func loadCards() ([]*Flashcard, Leitner, error) {
	cards := []*Flashcard{
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

	b1 := NewBox(0)
	b2 := NewBox(1)
	b3 := NewBox(2)
	b11 := NewBox(0)
	b22 := NewBox(1)
	b33 := NewBox(2)
	d1 := Deck{
		maxLevel: 3,
		boxes:    []Box{b1, b2, b3},
	}
	d2 := Deck{
		maxLevel: 3,
		boxes:    []Box{b11, b22, b33},
	}

	if err := d1.Insert(cards[0]); err != nil {
		return nil, Leitner{}, err
	}
	if err := d1.Insert(cards[1]); err != nil {
		return nil, Leitner{}, err
	}
	if err := d2.Insert(cards[2]); err != nil {
		return nil, Leitner{}, err
	}
	if err := d1.Insert(cards[3]); err != nil {
		return nil, Leitner{}, err
	}
	lt := NewLeitner(3, []Deck{d1, d2}, cooldownTest)

	return cards, lt, nil
}

func TestCoolDownPassed(t *testing.T) {
	cards, lt, err := loadCards()
	if err != nil {
		t.Errorf("unexpected error %v", err)
		return
	}

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
	if err != ErrCardsUnavailable {
		if err == nil {
			t.Errorf("wanted error but got nil, fc: %#+v", *fc)
		} else {
			t.Errorf("wanted another error, got: %v", err)
		}
		return
	}
}

func TestLeitner_Rate(t *testing.T) {
	cards, lt, err := loadCards()

	if err != nil {
		t.Errorf("unexpected error %v", err)
		return
	}

	// Deck0: 0:Zero, 1:One, 3:Three
	// Deck1: 2:Two

	t.Run("rating #1", func(t *testing.T) {
		err = lt.Rate(cards[3], 0, Bad)
		if err != nil {
			t.Errorf("unexpected error %v", err)
			return
		}

		if cards[3].L != 1 {
			t.Errorf("expected level 1, got %v", cards[3].L)
			return
		}

		if cards[3].Cd != CoolDownTest(2) {
			t.Errorf("expected cd 2, got %v", cards[3].Cd)
			return
		}

		if _, err := lt.decks[0].boxes[1].Get(3); err != nil {
			t.Errorf("Flashcard must be in box 1, but got error: %v", err)
			return
		}
	})

	t.Run("rating all", func(t *testing.T) {
		err = lt.Rate(cards[2], 1, Bad)
		if err != nil {
			t.Errorf("unexpected error %v", err)
			return
		}

		if cards[2].L != 0 {
			t.Errorf("expected level 1, got %v", cards[3].L)
			return
		}

		err = lt.Rate(cards[1], 0, Satisfactory)
		if err != nil {
			t.Errorf("unexpected error %v", err)
			return
		}

		if cards[1].L != 0 {
			t.Errorf("expected level 1, got %v", cards[3].L)
			return
		}

		err = lt.Rate(cards[0], 0, Good)
		if err != nil {
			t.Errorf("unexpected error %v", err)
			return
		}

		if cards[0].L != 1 {
			t.Errorf("expected level 1, got %v", cards[3].L)
			return
		}
	})

}

func TestCoolDown_TimeDynamic(t *testing.T) {
	tm = CoolDownTest(0)

	cards, lt, err := loadCards()
	if err != nil {
		t.Errorf("unexpected error %v", err)
		return
	}

	for _, c := range cards {
		c.CoolDown(cooldownTest)
	}

	fc, _, err := lt.GetRandom()
	if err != ErrCardsUnavailable {
		if err == nil {
			t.Errorf("wanted error but got nil, fc: %#+v", *fc)
		} else {
			t.Errorf("wanted another error, got: %v", err)
		}
		return
	}

	tm++
	fc, _, err = lt.GetRandom()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	fc.CoolDown(cooldownTest)

	fc, _, err = lt.GetRandom()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	fc.CoolDown(cooldownTest)

	fc, _, err = lt.GetRandom()
	if err != ErrCardsUnavailable {
		if err == nil {
			t.Errorf("wanted error but got nil, fc: %#+v", *fc)
		} else {
			t.Errorf("wanted another error, got: %v", err)
		}
		return
	}
}
