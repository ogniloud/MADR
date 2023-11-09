package models

type LoadDecksRequest struct {
	UserId UserId `json:"user_id"`
}

type LoadDecksResponse struct {
	Decks []DeckConfig `json:"decks"`
}

type GetFlashcardsByDeckIdRequest struct {
	DeckId DeckId `json:"deck_id"`
}

type GetFlashcardsByDeckIdResponse struct {
	Flashcards []Flashcard `json:"flashcards"`
}

type AddFlashcardToDeckRequest struct {
	UserId    UserId    `json:"user_id"`
	DeckId    DeckId    `json:"deck_id"`
	Flashcard Flashcard `json:"flashcard"`
}

type AddFlashcardToDeckResponse struct{}

type DeleteFlashcardFromDeckRequest struct {
	UserId      UserId      `json:"user_id"`
	FlashcardId FlashcardId `json:"flashcard_id"`
	DeckId      DeckId      `json:"deck_id"`
}

type DeleteFlashcardFromDeckResponse struct{}

type NewDeckWithFlashcardsRequest struct {
	UserId     UserId      `json:"user_id"`
	DeckConfig DeckConfig  `json:"deck_config"`
	Flashcards []Flashcard `json:"flashcards"`
}

type NewDeckWithFlashcardsResponse struct{}

type DeleteDeckRequest struct {
	UserId UserId `json:"user_id"`
	DeckId DeckId `json:"deck_id"`
}

type DeleteDeckResponse struct{}
