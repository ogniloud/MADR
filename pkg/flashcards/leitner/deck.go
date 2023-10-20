package leitner

import (
	"fmt"
	"math/rand"
)

var (
	ErrCardNotFound      = fmt.Errorf("flashcard not found")
	ErrCardAlreadyExists = fmt.Errorf("flashcard already exists")
	ErrCardsUnavailable  = fmt.Errorf("all flashcards unavailable")
	ErrBoxBadIndex       = fmt.Errorf("box index is bad")
	ErrDeckBadIndex      = fmt.Errorf("deck index is bad")
)

type BoxId = Level
type Box struct {
	L     Level
	Cards map[CardId]*Flashcard
}

func NewBox(l Level) Box {
	return Box{
		L:     l,
		Cards: map[CardId]*Flashcard{},
	}
}

func (b Box) Level() Level {
	return b.L
}

func (b Box) Get(id CardId) (*Flashcard, error) {
	fc, ok := b.Cards[id]
	if ok {
		return fc, nil
	}

	return nil, ErrCardNotFound
}

func (b Box) Delete(id CardId) error {
	_, ok := b.Cards[id]
	if ok {
		delete(b.Cards, id)
		return nil
	}

	return ErrCardNotFound
}

func (b Box) Add(flashcard *Flashcard) error {
	id := flashcard.Id

	_, ok := b.Cards[id]
	if ok {
		return ErrCardAlreadyExists
	}

	b.Cards[id] = flashcard
	return nil
}

// GetRandom returns a random AVAILABLE card from the box
// Randomization depends on implementation.
func (b Box) GetRandom() (*Flashcard, error) {
	for _, v := range b.Cards {
		if v.IsAvailable() {
			return v, nil
		}
	}
	return nil, ErrCardsUnavailable
}

type DeckId int
type Deck struct {
	MaxLevel Level
	Boxes    []Box
}

func NewDeck(l Level) Deck {
	d := Deck{
		MaxLevel: l,
		Boxes:    make([]Box, 0, l),
	}
	for i := 0; i < int(l); i++ {
		d.Boxes = append(d.Boxes, NewBox(l))
	}

	return d
}

// Insert inserts a flashcard to the box with
// corresponding Level.
func (d Deck) Insert(flashcard *Flashcard) error {
	l := flashcard.L

	b, err := d.Box(l)
	if err != nil {
		return err
	}

	err = b.Add(flashcard)
	return err
}

// Box returns a box by id. Returns an error if not exists
func (d Deck) Box(id BoxId) (Box, error) {
	if id < 0 || int(id) >= len(d.Boxes) {
		return Box{}, ErrBoxBadIndex
	}

	return d.Boxes[id], nil
}

func (d Deck) Delete(id CardId) error {
	for _, b := range d.Boxes {
		if err := b.Delete(id); err == nil {
			return nil
		}
	}
	return ErrCardNotFound
}

// GetRandom returns a random available flashcard from the random box.
// The random choose of a box depends on the level of the box.
func (d Deck) GetRandom(p []float32) (*Flashcard, error) {
	rd := rand.Float32()
	i := Level(0)
	for i != d.MaxLevel-1 && rd > p[i+1] {
		i++
	}

	b, _ := d.Box(i)
	fc, err := b.GetRandom()
	if err == nil {
		return fc, nil
	}

	for j := (i + 1) % d.MaxLevel; j != i; j = (j + 1) % d.MaxLevel {
		b, _ = d.Box(j)
		fc, err = b.GetRandom()
		if err == nil {
			return fc, nil
		}
	}
	return nil, ErrCardsUnavailable
}
