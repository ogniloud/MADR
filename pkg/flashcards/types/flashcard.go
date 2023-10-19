package types

// CardId is id of the flashcard. Must be unique.
type CardId int

// Word is id for flashcards
type Word string

// Level is a number in Leitner system where 1-Hottest, N-coldest
type Level int

// CoolDown is a timestamp after that a flashcard becomes available
type CoolDown interface {
	Passed() bool
}

// Backside is a field on a back of Flashcard with the Answer
// The handlers should use predefined types in flashcard package
// implementing Backside to handle answers correctly.
type Backside interface {

	// Answer returns a some item that defines the given Word.
	// Answer can be any object: definition, translation or music player.
	//
	// The implementation should process types of returned answer to use
	// them in the necessary goals.
	Answer() any
}

// Flashcard is a model of real flashcards where one side is a word
// and another one is an explanation
type Flashcard struct {
	// Id is id of a card that must be unique for each Deck
	Id CardId

	// Word is a top side of a card
	W Word

	B Backside

	// Level is a number of a Box in Leitner's system
	L Level

	// CoolDown is a timestamp before that the card is unavailable to learn
	Cd CoolDown
}

func (fc *Flashcard) IsAvailable() bool {
	return fc.Cd.Passed()
}

func (fc *Flashcard) CoolDown(f func(Level) CoolDown) {
	fc.Cd = f(fc.L)
}
