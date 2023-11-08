package handlers

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/ogniloud/madr/internal/flashcards/models"
	"github.com/ogniloud/madr/internal/flashcards/services/deck"
	httpjson "github.com/ogniloud/madr/internal/models"
)

type ErrorWriter interface {
	Error(w http.ResponseWriter, msg string, status int)
}

type Deck struct {
	s      deck.Service
	ew     ErrorWriter
	logger *log.Logger
}

func (d Deck) LoadDecks(w http.ResponseWriter, r *http.Request) {
	reqBody := models.LoadDecksRequest{}
	respBody := models.LoadDecksResponse{}

	if err := httpjson.FromJSON(&reqBody, r.Body); err != nil {
		d.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	decksMap, err := d.s.LoadDecks(reqBody.UserId)
	if err != nil {
		d.logger.Errorf("error while loading decks: %v", err)
		d.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.Decks = decksMap.Values()

	if err := httpjson.ToJSON(respBody, w); err != nil {
		d.logger.Errorf("error while writing response: %v", err)
		d.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
