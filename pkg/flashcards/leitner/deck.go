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
	l  Level
	av map[CardId]*Flashcard
}

func NewBox(l Level) Box {
	return Box{
		l:  l,
		av: map[CardId]*Flashcard{},
	}
}

func (b Box) Level() Level {
	return b.l
}

func (b Box) Get(id CardId) (*Flashcard, error) {
	fc, ok := b.av[id]
	if ok {
		return fc, nil
	}

	return nil, ErrCardNotFound
}

func (b Box) Delete(id CardId) error {
	_, ok := b.av[id]
	if ok {
		delete(b.av, id)
		return nil
	}

	return ErrCardNotFound
}

func (b Box) Add(flashcard *Flashcard) error {
	id := flashcard.Id

	_, ok := b.av[id]
	if ok {
		return ErrCardAlreadyExists
	}

	b.av[id] = flashcard
	return nil
}

// GetRandom returns a random AVAILABLE card from the box
// Randomization depends on implementation.
func (b Box) GetRandom() (*Flashcard, error) {
	for _, v := range b.av {
		if v.IsAvailable() {
			return v, nil
		}
	}
	return nil, ErrCardsUnavailable
}

type DeckId int
type Deck struct {
	maxLevel Level
	boxes    []Box
}

func NewDeck(l Level) Deck {
	d := Deck{
		maxLevel: l,
		boxes:    make([]Box, l),
	}

	for i := 0; i < int(l); i++ {
		d.boxes[i].l = Level(i)
		d.boxes[i].av = map[CardId]*Flashcard{}
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
	if id < 0 || int(id) >= len(d.boxes) {
		return Box{}, ErrBoxBadIndex
	}

	return d.boxes[id], nil
}

func (d Deck) Delete(id CardId) error {
	for _, b := range d.boxes {
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
	for i != d.maxLevel-1 && rd > p[i+1] {
		i++
	}

	b, _ := d.Box(i)
	fc, err := b.GetRandom()
	if err == nil {
		return fc, nil
	}

	for j := (i + 1) % d.maxLevel; j != i; j = (j + 1) % d.maxLevel {
		b, _ = d.Box(j)
		fc, err = b.GetRandom()
		if err == nil {
			return fc, nil
		}
	}
	return nil, ErrCardsUnavailable
}
