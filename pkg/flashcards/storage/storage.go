package storage

import (
	"time"
)

type DeckId = int
type UserId = int
type FlashcardId = int

type DeckConfig struct {
	Id      DeckId
	Name    string
	Private bool
}

type DeckConfigStorage interface {
	Get(id DeckId) (DeckConfig, error)
	Put(config DeckConfig) error
	Delete(id DeckId) error
}

type Word = string

// Backside is an abstract type for representation an answer
// of describing the Word. That can be just a string (translation, definition)
// or image, or sound - everything depends on Answer method.
type Backside interface {
	Answer() any
}

type Flashcard struct {
	Id     FlashcardId
	W      Word
	B      Backside
	DeckId DeckId
}

type FlashcardStorage interface {
	Get(id FlashcardId) (Flashcard, error)
	Put(flashcard Flashcard) error
	Delete(id FlashcardId) error
}

type LeitnerId = int
type Level = int

// CoolDown describes the timestamp.
// If the current time is less than the CoolDown state,
// the flashcard will not be shown for study.
type CoolDown struct {

	// the current simulation time.
	state time.Time

	// the function should return the time
	// that must pass from the moment this function is called
	// for the flashcard to become accessible again.
	f func(l Level) time.Duration
}

func NewCoolDown(startState time.Time, f func(l Level) time.Duration) CoolDown {
	return CoolDown{state: startState, f: f}
}

func (cd CoolDown) String() string {
	return cd.String()
}

// NextState updates the state of CoolDown relatively time.Now
func (cd CoolDown) NextState(l Level) {
	cd.state = time.Now().Add(cd.f(l))
}

// IsPassedNow returns true if state of CoolDown is not less than time.Now
func (cd CoolDown) IsPassedNow() bool {
	return cd.IsPassed(time.Now())
}

// IsPassed returns true if state of CoolDown is not less than t
func (cd CoolDown) IsPassed(t time.Time) bool {
	return cd.state.Compare(t) != -1
}

type UserLeitner struct {
	Id          LeitnerId
	UserId      UserId
	FlashcardId FlashcardId
	Level       Level
	CoolDown    CoolDown
}

type LeitnerStorage interface {
	Get(id LeitnerId) (UserLeitner, error)
	Put(leitner UserLeitner) error
	Delete(id LeitnerId) error
}
