package models

type LoadDecksRequest struct {
	UserId UserId `json:"user_id"`
}

type LoadDecksResponse struct {
	Decks []DeckConfig `json:"decks"`
}
