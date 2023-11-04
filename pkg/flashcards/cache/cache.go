// Package cache provides simple access to in-memory cache.
// The major thing for using this is accessing user's decks
package cache

import (
	"sync"
)

// Cache implements thread-safe Load, Store and Delete functions.
type Cache struct {
	*sync.Map
}

func New() Cache {
	return Cache{&sync.Map{}}
}
