package models

// Mark is a type of rating marks of flashcards in Leitner's system.
type Mark int

// Matching is used for exercises connected with matching
// words with their definitions or other answers.
type Matching struct {
	M     map[string]string    `json:"pairs"`
	Cards map[string]Flashcard `json:"cards"`
}

// Text is used for exercises connected with insertion
// words in a sentence.
type Text struct {
	// The words to guess are escaped like this: <div style=\"display:none;\"> WORD </div>
	T string `json:"text"`

	// Options for filling the spaces
	Opts  []string             `json:"options"`
	Cards map[string]Flashcard `json:"cards"`
}
