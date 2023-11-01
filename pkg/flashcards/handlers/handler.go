package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ogniloud/madr/pkg/flashcards/models"
	"github.com/ogniloud/madr/pkg/flashcards/services/deck"
)

type ErrorWriter interface {
	Error(w http.ResponseWriter, msg string, status int)
}

type Deck struct {
	s  deck.Service
	ew ErrorWriter
}

func (d Deck) LoadDecks(w http.ResponseWriter, r *http.Request) {
	reqBody := models.LoadDecksRequest{}
	respBody := models.LoadDecksResponse{}

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&reqBody); err != nil {
		d.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	decksMap, err := d.s.LoadDecks(reqBody.UserId)
	if err != nil {
		d.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.Decks = decksMap.Values()

	enc := json.NewEncoder(w)
	if err := enc.Encode(respBody); err != nil {
		d.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
