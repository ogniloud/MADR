package models

// LoadDecksRequest is a struct that defines the request body for
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

// LoadDecksResponse is a struct that defines the response body for
// loading deck handler.
type LoadDecksResponse struct {
	// Decks is a slice of deck with their configs.
	//
	// Extensions:
	// ---
	// x-property-value: value
	// x-property-array:
	//   - value1
	//   - value2
	// x-property-array-obj:
	//   - name: obj
	//     value: field
	// ---
	Decks []DeckConfig `json:"decks"`
}

// GetFlashcardsByDeckIdRequest is a struct that defines the request body for
// loading cards from the deck handler.
//
//swagger:model getFlashcardsByDeckIdRequest
type GetFlashcardsByDeckIdRequest struct {
	// DeckId defines the deck which flashcards will be loaded from.
	//
	// required: true
	// example: 123
	DeckId DeckId `json:"deck_id"`
}

// GetFlashcardsByDeckIdResponse is a struct that defines the response body for
// loading cards from the deck handler.
type GetFlashcardsByDeckIdResponse struct {
	// Flashcards is a slice of flashcards that is returned to the client.
	//
	// Extensions:
	// ---
	// x-property-value: value
	// x-property-array:
	//   - value1
	//   - value2
	// x-property-array-obj:
	//   - name: obj
	//     value: field
	// ---
	Flashcards []Flashcard `json:"flashcards"`
}

// AddFlashcardToDeckRequest is a struct that defines the request body for
// adding cards to the deck handler.
//
//swagger:model addFlashcardToDeckRequest
type AddFlashcardToDeckRequest struct {
	// UserId is an ID of the user in a storage.
	//
	// required: true
	// example: 189
	UserId UserId `json:"user_id"`

	// DeckId is an ID of the deck which the flashcard will be added to.
	//
	// required: true
	// example: 123
	DeckId DeckId `json:"deck_id"`

	// Word is a head side of the card.
	//
	// required: true
	// example: Aboba
	Word Word `json:"word"`

	// Backside is a back side of the card.
	//
	// required: true
	// example: Backside{Type: TypeDefinition, Value: "President"}
	Backside Backside `json:"backside"`

	// Answer is an answer for exercises.
	//
	// required: true
	// example: "aboba"
	Answer Answer `json:"answer"`
}

// AddFlashcardToDeckResponse is a struct that defines the response body for
// adding cards to the deck handler.
type AddFlashcardToDeckResponse struct{}

// DeleteFlashcardFromDeckRequest is a struct that defines the request body for
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

// DeleteFlashcardFromDeckResponse is a struct that defines the response body for
// deleting cards from the deck handler.
type DeleteFlashcardFromDeckResponse struct{}

// NewDeckWithFlashcardsRequest is a struct that defines the request body for
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
	// example: Aboba123
	Name string `json:"name"`

	// Flashcards is a slice of flashcards that will be added.
	//
	// Extensions:
	// ---
	// x-property-value: value
	// x-property-array:
	//   - value1
	//   - value2
	// x-property-array-obj:
	//   - name: obj
	//     value: field
	// ---
	Flashcards []struct {
		Word     Word     `json:"word"`
		Backside Backside `json:"backside"`
		Answer   Answer   `json:"answer"`
	} `json:"flashcards"`
}

// NewDeckWithFlashcardsResponse is a struct that defines the response body for
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

	// DeckId is an ID of the deck which will be deleted.
	//
	// required: true
	// example: 123
	DeckId DeckId `json:"deck_id"`
}

// DeleteDeckResponse is a struct that defines the response body for the
// deleting the deck handler.
type DeleteDeckResponse struct{}

// RandomCardDeckRequest is a struct that defines the request body for the
// getting a random card from the deck.
//
//swagger:model randomCardDeckRequest
type RandomCardDeckRequest struct {
	// UserId is an ID of the user in a storage.
	//
	// required: true
	// example: 189
	UserId UserId `json:"user_id"`

	// DeckId is an ID of the deck where the flashcard are taken from.
	//
	// required: true
	// example: 123
	DeckId DeckId `json:"deck_id"`
}

// RandomCardDeckResponse is a struct that defines the response body for
// getting a random card from the deck.
type RandomCardDeckResponse struct {
	// Flashcard is a flashcard that is returned to the user.
	Flashcard Flashcard `json:"flashcard"`
}

// RandomNCardsDeckRequest is a struct that defines the request body for
// getting n random cards from the deck.
//
//swagger:model randomNCardsDeckRequest
type RandomNCardsDeckRequest struct {
	// UserId is an ID of the user in a storage.
	//
	// required: true
	// example: 189
	UserId UserId `json:"user_id"`

	// DeckId is an ID of the deck where the flashcard are taken from.
	//
	// required: true
	// example: 123
	DeckId DeckId `json:"deck_id"`

	// Amount of cards to get
	N int `json:"n"`
}

// RandomNCardsDeckResponse is a struct that defines the response body for
// getting n random cards from the deck.
type RandomNCardsDeckResponse struct {
	// Flashcard is a flashcard that is returned to the user.
	Flashcards []Flashcard `json:"flashcards"`
}

// RandomCardRequest is a struct that defines the request body for
// getting a random card from all the decks.
//
//swagger:model randomCardRequest
type RandomCardRequest struct {
	// UserId is an ID of the user in a storage.
	//
	// required: true
	// example: 189
	UserId UserId `json:"user_id"`
}

// RandomCardResponse is a struct that defines the response body for
// getting a random card from all the decks.
type RandomCardResponse struct {
	// Flashcard is a flashcard that is returned to the user.
	Flashcard Flashcard `json:"flashcard"`
}

// RandomNCardsRequest is a struct that defines the request body for
// getting n random cards from all the decks.
//
//swagger:model randomNCardsRequest
type RandomNCardsRequest struct {
	// UserId is an ID of the user in a storage.
	//
	// required: true
	// example: 189
	UserId UserId `json:"user_id"`

	// Amount of cards to get
	N int `json:"n"`
}

// RandomNCardsResponse is a struct that defines the response body for
// getting n random cards from all the decks.
type RandomNCardsResponse struct {
	// Flashcard is a flashcard that is returned to the user.
	Flashcards []Flashcard `json:"flashcards"`
}

// RateRequest is a struct that defines the request body for the
// rating a card.
//
//swagger:model rateRequest
type RateRequest struct {
	// UserId is an ID of the user in a storage.
	//
	// required: true
	// example: 189
	UserId UserId `json:"user_id"`

	// FlashcardId is an ID of the flashcard.
	//
	// required: true
	// example: 1111
	FlashcardId FlashcardId `json:"flashcard_id"`

	// Mark is a mark for the card.
	//
	// required: true
	// example: 0, 1 or 2
	Mark Mark `json:"mark"`
}

// RateResponse is a struct that defines the request body for the
// rating a card.
type RateResponse struct{}

// RandomMatchingRequest is a struct that defines the request body for
// getting a random matching exercise from all the decks.
//
//swagger:model randomMatchingRequest
type RandomMatchingRequest struct {
	// UserId is an ID of the user in a storage.
	//
	// required: true
	// example: 189
	UserId UserId `json:"user_id"`

	// Size is the max amount of cards to include int the exercise
	//
	// required: true
	// example: 5
	Size int `json:"size"`
}

// RandomMatchingRequest is a struct that defines the request body for
// getting a random matching from all the decks.
type RandomMatchingResponse struct {
	Matching Matching `json:"matching"`
}

// RandomMatchingDeckRequest is a struct that defines the request body for
// getting a random matching exercise from the deck.
//
//swagger:model randomMatchingDeckRequest
type RandomMatchingDeckRequest struct {
	// UserId is an ID of the user in a storage.
	//
	// required: true
	// example: 189
	UserId UserId `json:"user_id"`

	// DeckId is an ID of the deck where the flashcard are taken from.
	//
	// required: true
	// example: 123
	DeckId DeckId `json:"deck_id"`

	// Size is the max amount of cards to include int the exercise
	//
	// required: true
	// example: 5
	Size int `json:"size"`
}

// RandomMatchingDeckRequest is a struct that defines the request body for
// getting a random matching from the deck.
type RandomMatchingDeckResponse struct {
	Matching Matching `json:"matching"`
}

// RandomTextRequest is a struct that defines the request body for
// getting a random text exercise from all the decks.
//
//swagger:model randomTextRequest
type RandomTextRequest struct {
	// UserId is an ID of the user in a storage.
	//
	// required: true
	// example: 189
	UserId UserId `json:"user_id"`

	// Size is the max amount of cards to include int the exercise
	//
	// required: true
	// example: 5
	Size int `json:"size"`
}

// RandomTextResponse is a struct that defines the response body for
// getting a random text exercise from all the decks.
type RandomTextResponse struct {
	Text Text `json:"text"`
}

// RandomTextDeckRequest is a struct that defines the request body for
// getting a random text exercise from all the decks.
//
//swagger:model randomTextDeckRequest
type RandomTextDeckRequest struct {
	// UserId is an ID of the user in a storage.
	//
	// required: true
	// example: 189
	UserId UserId `json:"user_id"`

	// DeckId is an ID of the deck where the flashcard are taken from.
	//
	// required: true
	// example: 123
	DeckId DeckId `json:"deck_id"`

	// Size is the max amount of cards to include int the exercise
	//
	// required: true
	// example: 5
	Size int `json:"size"`
}

// RandomTextDeckResponse is a struct that defines the response body for
// getting a random text exercise from all the decks.
type RandomTextDeckResponse struct {
	Text Text `json:"text"`
}

// GetFlashcardByIdResponse is a response that contains a flashcard.
type GetFlashcardByIdResponse struct {
	Flashcard Flashcard `json:"flashcard"`
}

type AppendBacksideRequest struct {
	UserId    UserId      `json:"user_id"`
	CardId    FlashcardId `json:"card_id"`
	Backsides []Backside  `json:"backsides"`
}

type AppendBacksideResponse struct{}
