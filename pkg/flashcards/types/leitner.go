package types

type LeitnerId int

type Rate int

const (
	Bad = Rate(iota)
	Satisfactry
	Good
)

// Leitner is an abstract data structure consisting of
// Boxes. Each box has a temperature Level. It means that
// the hotter the box the higher chance to be chosen by GetRandom.
// Leitner should consist of deck (deck of decks).
//
// Also, for each flashcard defined CoolDown. A flashcard can't be returned
// if CoolDown is not passed.
type Leitner interface {
	Decks
	GetRandom() (*Flashcard, DeckId, error)
	Rate(*Flashcard, DeckId, Rate) error
}
