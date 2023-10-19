package types

type LeitnerId int

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
type Leitner interface {
	Decks

	// GetRandom means take a random card from a random deck
	GetRandom() (*Flashcard, DeckId, error)

	// Rate takes a mark from the user and inserts the card
	// in a corresponding to its level box
	Rate(*Flashcard, DeckId, Rate) error
}
