package leitner

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

type Translation string

func (t Translation) Answer() any {
	return t
}
