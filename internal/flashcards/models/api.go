package models

// LoadDecksRequest is a struct that defines the request body for the
// loading deck handler.
//
//swagger:model loadDecksRequest
type LoadDecksRequest struct {
	// UserId is used for loading for the cards that the user has.
	//
	// required: true
	// example: 189
	UserId UserId `json:"user_id"`
}

// LoadDecksResponse is a struct that defines the request body for the
// loading deck handler.
type LoadDecksResponse struct {
	// Decks is a slice of deck with their configs.
	Decks []DeckConfig `json:"decks"`
}

// GetFlashcardsByDeckIdRequest is a struct that defines the request body for the
// loading cards from the deck handler.
//
//swagger:model getFlashcardsByDeckIdRequest
type GetFlashcardsByDeckIdRequest struct {
	// DeckId defines the deck which flashcards will be loaded.
	//
	// required: true
	// example: 123
	DeckId DeckId `json:"deck_id"`
}

// GetFlashcardsByDeckIdResponse is a struct that defines the request body for the
// loading cards from the deck handler.
type GetFlashcardsByDeckIdResponse struct {
	// Flashcards is a slice of flashcards that is returned to the client.
	Flashcards []Flashcard `json:"flashcards"`
}

// AddFlashcardToDeckRequest is a struct that defines the request body for the
// adding cards to the deck handler.
//
//swagger:model addFlashcardToDeckRequest
type AddFlashcardToDeckRequest struct {
	// UserId is an ID of the user in a storage.
	//
	// required: true
	// example: 189
	UserId UserId `json:"user_id"`

	// DeckId is an ID of the deck where the flashcard will be added.
	//
	// required: true
	// example: 123
	DeckId DeckId `json:"deck_id"`

	// Word is a head side of the card.
	//
	// required: true
	// example: "Aboba"
	Word Word `json:"word"`

	// Backside is a back side of the card.
	//
	// required: true
	// example: Backside{Type: TypeDefinition, Value: "President"}
	Backside Backside `json:"backside"`
}

// AddFlashcardToDeckResponse is a struct that defines the request body for the
// adding cards to the deck handler.
type AddFlashcardToDeckResponse struct{}

// DeleteFlashcardFromDeckRequest is a struct that defines the request body for the
// deleting cards from the deck handler.
//
//swagger:model deleteFlashcardFromDeckRequest
type DeleteFlashcardFromDeckRequest struct {
	// UserId is an ID of the user.
	//
	// required: true
	// example: 189
	UserId UserId `json:"user_id"`

	// FlashcardId is an ID of the flashcard.
	//
	// required: true
	// example: 666
	FlashcardId FlashcardId `json:"flashcard_id"`
}

// DeleteFlashcardFromDeckResponse is a struct that defines the request body for the
// deleting cards from the deck handler.
type DeleteFlashcardFromDeckResponse struct{}

// NewDeckWithFlashcardsRequest is a struct that defines the request body for the
// creating a new deck handler.
//
//swagger:model newDeckWithFlashcardsRequest
type NewDeckWithFlashcardsRequest struct {
	// UserId is an ID of the user in a storage.
	//
	// required: true
	// example: 189
	UserId UserId `json:"user_id"`

	// Name is a name of the deck
	//
	// required: true
	// example: "Aboba123"
	Name string `json:"name"`

	// Flashcards is a slice of flashcards that will be added.
	Flashcards []Flashcard `json:"flashcards"`
}

// NewDeckWithFlashcardsResponse is a struct that defines the request body for the
// creating a new deck handler.
type NewDeckWithFlashcardsResponse struct{}

// DeleteDeckRequest is a struct that defines the request body for the
// deleting the deck handler.
//
//swagger:model deleteDeckRequest
type DeleteDeckRequest struct {
	// UserId is an ID of the user in a storage.
	//
	// required: true
	// example: 189
	UserId UserId `json:"user_id"`

	// DeckId is an ID of the deck where the flashcard will be added.
	//
	// required: true
	// example: 123
	DeckId DeckId `json:"deck_id"`
}

// DeleteDeckResponse is a struct that defines the request body for the
// deleting the deck handler.
type DeleteDeckResponse struct{}
