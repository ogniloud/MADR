package leitner

import (
	"time"
)

// Flashcard is a model of real flashcards where one side is a word
// and another one is an explanation
type Flashcard struct {
	// Id is id of a card that must be unique for each Deck
	Id CardId

	// Word is a top side of a card
	W Word

	B Backside

	// Level is a number of a Box in Leitner's system
	L Level

	// CoolDown is a timestamp before that the card is unavailable to learn
	Cd CoolDown
}

// CardId is id of the flashcard. Must be unique.
type CardId int

// Word is id for flashcards
type Word string

// Level is a number in Leitner system where 1-Hottest, N-coldest
type Level int

// CoolDown is a timestamp after that a flashcard becomes available
type CoolDown interface {
	Passed() bool
}

func (fc *Flashcard) IsAvailable() bool {
	return fc.Cd.Passed()
}

func (fc *Flashcard) CoolDown(f func(Level) CoolDown) {
	fc.Cd = f(fc.L)
}

// CoolDownTime is a cool down relatively real time
type CoolDownTime time.Time

// Passed if timestamp of Now is later than cool down timestamp
func (cd CoolDownTime) Passed() bool {
	return time.Now().After(time.Time(cd))
}

// CoolDownFunc returns new cool down for each level
func CoolDownFunc(l Level) CoolDownTime {
	return CoolDownTime(time.Now().Add(time.Duration((l*l+1)*24) * time.Hour))
}
