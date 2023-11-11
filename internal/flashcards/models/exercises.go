package models

type Matching struct {
	M     map[string]string
	Cards map[string]Flashcard
}

type Text struct {
	T     string
	Opts  []string
	Cards map[string]Flashcard
}
