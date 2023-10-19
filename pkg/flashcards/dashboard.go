package flashcards

import "github.com/ogniloud/madr/pkg/flashcards/types"

type Service interface {
	GetLeitnerById(id types.LeitnerId)
}
