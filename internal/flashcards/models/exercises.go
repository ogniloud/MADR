package models

// Mark is a type of rating marks of flashcards in Leitner's system.
type Mark int

// Matching is used for exercises connected with matching
// words with their definitions or other answers.
type Matching struct {
	M     map[string]string
	Cards map[string]Flashcard
}

// Text is used for exercises connected with insertion
// words in a sentence.
type Text struct {
	T     string
	Opts  []string
	Cards map[string]Flashcard
}
