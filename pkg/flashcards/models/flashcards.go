package models

import "time"

type (
	DeckId      = int
	UserId      = int
	FlashcardId = int
)

type DeckConfig struct {
	DeckId DeckId `json:"deck_id"`
	UserId UserId `json:"user_id"`
	Name   string `json:"name"`
}

type UserInfo struct {
	Id     UserId `json:"user_id"`
	MaxBox Box    `json:"max_box"`
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

func (cd CoolDown) String() string {
	return cd.State.String()
}

// NextState updates the state of CoolDown relatively f
func (cd CoolDown) NextState(b Box, f func(Box) time.Time) {
	cd.State = f(b)
}

// IsPassedNow returns true if state of CoolDown is not less than time.Now
func (cd CoolDown) IsPassedNow() bool {
	return cd.IsPassed(CoolDown{State: time.Now()})
}

// IsPassed returns true if state of CoolDown is not less than t
func (cd CoolDown) IsPassed(t CoolDown) bool {
	return cd.State.Compare(t.State) == -1
}

type UserLeitner struct {
	Id          LeitnerId   `json:"id"`
	UserId      UserId      `json:"user_id"`
	FlashcardId FlashcardId `json:"card_id"`
	Box         Box         `json:"level"`
	CoolDown    CoolDown    `json:"cooldown"`
}

type Decks map[DeckId]DeckConfig

func (d Decks) Keys() []DeckId {
	ids := make([]DeckId, 0, len(d))
	for k := range d {
		ids = append(ids, k)
	}

	return ids
}
