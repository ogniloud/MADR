package storage

import "github.com/ogniloud/madr/pkg/flashcards/leitner"

type UserId int

type Leitner struct {
	LtId   leitner.Id
	UserId UserId
	DkId   leitner.DeckId
	Decks  []leitner.Level
	Level  leitner.Level
}

type Flashcard struct {
	W  leitner.Word
	B  leitner.Backside
	Cd leitner.CoolDownTime
}

type Deck struct {
	DkId leitner.DeckId
	Fcs  []Flashcard
}

type LeitnerStorer interface {
	Get(leitner.Id) (Leitner, error)
	GetAllByUserId(UserId) ([]Leitner, error)
	Put(Leitner) error
	Update(Leitner) error
}

type DeckStorer interface {
	Get(leitner.DeckId) (Deck, error)
	Put(Deck) error
	Update(Deck) error
}
