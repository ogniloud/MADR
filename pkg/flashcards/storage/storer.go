package storage

import "github.com/ogniloud/madr/pkg/flashcards/leitner"

type UserConfig struct {
	Id       int
	MaxLevel leitner.Level
}

type UserConfigStorer interface {
	Store(deck UserConfig) error
	Get(id int) (UserConfig, error)
}

type Card struct {
	Id       leitner.CardId
	DeckId   leitner.DeckId
	Word     leitner.Word
	Backside leitner.Backside
}

type CardStorer interface {
	Store(card Card) error
	Get(id leitner.CardId) (Card, error)
}

type DeckPrivacy struct {
	DeckId  leitner.DeckId
	Private bool
}

type DeckPrivacyStorer interface {
	Store(deck DeckPrivacy) error
	Get(id leitner.DeckId) (DeckPrivacy, error)
}

type Leitner struct {
	Id     leitner.Id
	UserId int
	CardId leitner.CardId
	Level  leitner.Level
	Cd     leitner.CoolDown
}

type LeitnerStorer interface {
	Store(l Leitner) error
	Get(leitner.Id) (Leitner, error)
	GetAllByUserId(int) ([]Leitner, error)
}
