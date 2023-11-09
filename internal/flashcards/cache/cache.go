// Package cache provides simple access to in-memory cache.
// The major thing for using this is accessing user's decks
package cache

import (
	"fmt"
	"sync"
)

var ErrNotFound = fmt.Errorf("value not found")

// Cache implements thread-safe Load, Store functions.
type Cache struct {
	m *sync.Map
}

func New() Cache {
	return Cache{&sync.Map{}}
}

// Load returns a value if stored in cache.
func (c Cache) Load(k any) (any, error) {
	v, ok := c.m.Load(k)
	if !ok {
		return nil, ErrNotFound
	}

	return v, nil
}

// Store puts the value v with key k.
func (c Cache) Store(k, v any) error {
	c.m.Store(k, v)
	return nil
}

func (c Cache) Delete(k any) error {
	c.m.Delete(k)
	return nil
}
