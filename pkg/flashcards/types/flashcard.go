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

type Flashcard struct {
	Id CardId
	W  Word
	B  Backside
	L  Level
	Cd CoolDown
}

func (fc *Flashcard) IsAvailable() bool {
	return fc.Cd.Passed()
}

func (fc *Flashcard) CoolDown(f func(Level) CoolDown) {
	fc.Cd = f(fc.L)
}
