package models

import (
	"time"
)

type (
	DeckId      = int
	UserId      = int
	FlashcardId = int
)

// DeckConfig contains information about a particular deck.
// DeckId is not a primary key, for each user a configuration can be different.
type DeckConfig struct {
	DeckId DeckId `json:"deck_id"`

	// UserId means that a user with UserId has the deck with id DeckId.
	UserId UserId `json:"user_id"`

	// Name is a name of the deck which the user assigned to it.
	Name string `json:"name"`
}

// UserInfo contains information about the user.
type UserInfo struct {
	UserId UserId `json:"user_id"` // primary

	// MaxBox is a maximal amount of boxes in Leitner's system.
	MaxBox Box `json:"max_box"`
}

type (
	Word   = string
	Answer = string

	// BacksideType tells how to imagine a value of Backside.
	BacksideType int
)

const (
	Translation = BacksideType(iota)
	Definition
)

// Backside is an abstract type for representation an answer
// of describing the Word. That can be just a string (translation, definition)
// or image, or sound - everything depends on Answer method.
type Backside struct {
	Type  BacksideType `json:"type"`
	Value string       `json:"value"`
}

// ParseBackside returns a new string from value of backside if needed.
//
// For example, from base64 to another string.
func ParseBackside(b Backside) string {
	switch b.Type {
	case Translation, Definition:
		return b.Value
	default:
		return b.Value
	}
}

// Flashcard is a model of real flashcards with front side with word
// and back side containing some information that describes a word.
type Flashcard struct {
	Id FlashcardId `json:"id"`
	W  Word        `json:"word"`
	A  Answer      `json:"answer"`
	B  Backside    `json:"backside"`

	// DeckId shows which deck the flashcard belongs to.
	DeckId DeckId `json:"deck_id"`
}

type (
	LeitnerId = int
	Box       = int
)

// CoolDown describes the timestamp.
// If the current time is less than the CoolDown state,
// the flashcard will not be shown for study.
type CoolDown time.Time

// NextState updates the state of cd relatively f.
func (cd *CoolDown) NextState(b Box, f func(Box) time.Time) {
	*cd = CoolDown(f(b))
}

// IsPassedNow returns true if state of CoolDown is not less than time.Now.
func (cd *CoolDown) IsPassedNow() bool {
	return cd.IsPassed(CoolDown(time.Now()))
}

// IsPassed returns true if state of cd is not less than t.
func (cd *CoolDown) IsPassed(t CoolDown) bool {
	t1 := time.Time(*cd)
	return t1.Before(time.Time(t))
}

// UserLeitner is a configuration of Leitner's system for each user.
//
// For each user the structure contains a FlashcardId with its Box and CoolDown.
// CoolDown is a timestamp after which the card can be examined again.
//
// There is a bijection between (UserId, FlashcardId) and UserLeitner.
type UserLeitner struct {
	Id          LeitnerId   `json:"id"` // primary
	UserId      UserId      `json:"user_id"`
	FlashcardId FlashcardId `json:"card_id"`

	// Box is an item in Leitner's system showing frequency
	// of studying flashcards. The higher box, the more cool down.
	Box Box `json:"box"`

	// CoolDown is a timestamp showing when the flashcard will be ready
	// to be examined.
	CoolDown CoolDown `json:"cool_down"`
}

// Decks is a map of decks which config can be got by id.
type Decks map[DeckId]DeckConfig

func (d Decks) Keys() []DeckId {
	ids := make([]DeckId, 0, len(d))
	for k := range d {
		ids = append(ids, k)
	}

	return ids
}

func (d Decks) Values() []DeckConfig {
	ids := make([]DeckConfig, 0, len(d))
	for _, v := range d {
		ids = append(ids, v)
	}

	return ids
}
