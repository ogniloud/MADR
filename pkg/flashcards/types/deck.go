package types

type BoxId Level

type Box interface {
	// Level return a level of Box (0-Hot, Max-Cold)
	Level() Level

	Get(CardId) (*Flashcard, error)
	Delete(CardId) error
	Add(*Flashcard) error

	// Update - looking up by CardId.
	// Level must be the same.
	Update(*Flashcard) error

	// GetRandom returns a random AVAILABLE card from the box
	// Randomization depends on implementation.
	GetRandom() (*Flashcard, error)
}

type Boxes interface {

	// Insert inserts a flashcard to the box with
	// corresponding Level.
	Insert(*Flashcard) error

	// Box returns a box by id. Returns an error if not exists
	Box(BoxId) (Box, error)
}

type DeckId int

type Deck interface {
	Boxes

	// GetRandom returns a random available flashcard from the random box.
	// The random choose of a box depends on the level of the box.
	GetRandom() (*Flashcard, error)
}

type Decks interface {
	Deck(DeckId) (Deck, error)
}
