package storage

import (
	"time"
)

type (
	DeckId      = int
	UserId      = int
	FlashcardId = int
)

type DeckConfig struct {
	Id   DeckId `json:"id"`
	Name string `json:"name"`
}

type (
	Word         = string
	BacksideType int
)

// Backside is an abstract type for representation an answer
// of describing the Word. That can be just a string (translation, definition)
// or image, or sound - everything depends on Answer method.
type Backside struct {
	Type  BacksideType `json:"type"`
	Value string       `json:"value"`
}

type Flashcard struct {
	Id     FlashcardId `json:"id"`
	W      Word        `json:"word"`
	B      Backside    `json:"backside"`
	DeckId DeckId      `json:"deck_id"`
}

type (
	LeitnerId = int
	Box       = int
)

// CoolDown describes the timestamp.
// If the current time is less than the CoolDown state,
// the flashcard will not be shown for study.
type CoolDown struct {
	// the current simulation time.
	State time.Time `json:"state"`
}

func NewCoolDown(startState time.Time) CoolDown {
	return CoolDown{State: startState}
}

func (cd CoolDown) String() string {
	return cd.State.String()
}

// NextState updates the state of CoolDown relatively f
func (cd CoolDown) NextState(b Box, f func(Box) time.Time) {
	cd.State = f(b)
}

// IsPassedNow returns true if state of CoolDown is not less than time.Now
func (cd CoolDown) IsPassedNow() bool {
	return cd.IsPassed(time.Now())
}

// IsPassed returns true if state of CoolDown is not less than t
func (cd CoolDown) IsPassed(t time.Time) bool {
	return cd.State.Compare(t) != -1
}

type UserLeitner struct {
	Id          LeitnerId   `json:"id"`
	UserId      UserId      `json:"user_id"`
	FlashcardId FlashcardId `json:"card_id"`
	Box         Box         `json:"level"`
	CoolDown    CoolDown    `json:"cooldown"`
}

type Decks map[DeckId]DeckConfig

//go:generate go run github.com/vektra/mockery/v2@v2.36.0 --name Storage
type Storage interface {
	GetDecksByUserId(id UserId) (Decks, error)
	GetFlashcardsByDeckId(id DeckId) ([]Flashcard, error)
	PutFlashcard(id DeckId, card Flashcard) error
	PutAllFlashcards(id DeckId, cards []Flashcard) error
	DeleteFlashcardFromDeck(id DeckId, cardId FlashcardId) error
	PutNewDeck(config DeckConfig) error
	DeleteDeck(id DeckId) error
}
