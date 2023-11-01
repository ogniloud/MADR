package cache

import (
	"sync"

	"github.com/ogniloud/madr/pkg/flashcards/models"
)

type Cache = *sync.Map

type CachedRandom []models.FlashcardId
