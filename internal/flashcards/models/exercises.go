package models

type Matching struct {
	M     map[string]string
	Cards map[string]Flashcard
}

type Sentence struct {
	S     string
	Opts  []string
	Cards map[string]Flashcard
}
