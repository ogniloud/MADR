package leitner

import (
	"fmt"
	ftypes "github.com/ogniloud/madr/pkg/flashcards/types"
	"math/rand"
)

var (
	ErrCardNotFound      = fmt.Errorf("flashcard not found")
	ErrCardAlreadyExists = fmt.Errorf("flashcard already exists")
	ErrCardsUnavailable  = fmt.Errorf("all flashcards unavailable")
	ErrBoxBadIndex       = fmt.Errorf("box index is bad")
	ErrDeckBadIndex      = fmt.Errorf("deck index is bad")
)

type Box struct {
	l  ftypes.Level
	av map[ftypes.CardId]*ftypes.Flashcard
}

func NewBox(l ftypes.Level) Box {
	return Box{
		l:  l,
		av: map[ftypes.CardId]*ftypes.Flashcard{},
	}
}

func (b Box) Level() ftypes.Level {
	return b.l
}

func (b Box) Get(id ftypes.CardId) (*ftypes.Flashcard, error) {
	fc, ok := b.av[id]
	if ok {
		return fc, nil
	}

	return nil, ErrCardNotFound
}

func (b Box) Delete(id ftypes.CardId) error {
	_, ok := b.av[id]
	if ok {
		delete(b.av, id)
		return nil
	}

	return ErrCardNotFound
}

func (b Box) Add(flashcard *ftypes.Flashcard) error {
	id := flashcard.Id

	_, ok := b.av[id]
	if ok {
		return ErrCardAlreadyExists
	}

	b.av[id] = flashcard
	return nil
}

func (b Box) GetRandom() (*ftypes.Flashcard, error) {
	for _, v := range b.av {
		if v.IsAvailable() {
			return v, nil
		}
	}
	return nil, ErrCardsUnavailable
}

type Deck struct {
	maxLevel ftypes.Level
	boxes    []ftypes.Box
}

func (d Deck) Insert(flashcard *ftypes.Flashcard) error {
	l := flashcard.L

	b, err := d.Box(l)
	if err != nil {
		return err
	}

	err = b.Add(flashcard)
	return err
}

func (d Deck) Box(id ftypes.Level) (ftypes.Box, error) {
	if id < 0 || int(id) >= len(d.boxes) {
		return nil, ErrBoxBadIndex
	}

	return d.boxes[id], nil
}

func (d Deck) Delete(id ftypes.CardId) error {
	for _, b := range d.boxes {
		if err := b.Delete(id); err == nil {
			return nil
		}
	}
	return ErrCardNotFound
}

func (d Deck) GetRandom(p []float32) (*ftypes.Flashcard, error) {
	rd := rand.Float32()
	i := ftypes.Level(0)
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
