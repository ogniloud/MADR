package handlers

import (
	"net/http"
	"time"

	"github.com/charmbracelet/log"
	"github.com/ogniloud/madr/internal/flashcards/models"
	"github.com/ogniloud/madr/internal/flashcards/services/deck"
	"github.com/ogniloud/madr/internal/ioutil"
)

type Deck struct {
	s      deck.IService
	ew     ioutil.ErrorWriter
	logger *log.Logger
}

func (d Deck) LoadDecks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		d.ew.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}

	reqBody := models.LoadDecksRequest{}
	respBody := models.LoadDecksResponse{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
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

	if err := ioutil.ToJSON(respBody, w); err != nil {
		d.logger.Errorf("error while writing response: %v", err)
		d.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (d Deck) GetFlashcardsByDeckId(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		d.ew.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}

	reqBody := models.GetFlashcardsByDeckIdRequest{}
	respBody := models.GetFlashcardsByDeckIdResponse{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		d.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ids, err := d.s.GetFlashcardsIdByDeckId(reqBody.DeckId)
	if err != nil {
		d.logger.Errorf("error while loading ids of cards: %v", err)
		d.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.Flashcards = make([]models.Flashcard, len(ids))

	// should be replaced with one read transaction
	for i := 0; i < len(ids); i++ {
		card, err := d.s.GetFlashcardById(ids[i])
		if err != nil {
			d.logger.Errorf("error while loading a card: %v", err)
			d.ew.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		respBody.Flashcards[i] = card
	}

	if err := ioutil.ToJSON(respBody, w); err != nil {
		d.logger.Errorf("error while writing response: %v", err)
		d.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (d Deck) AddFlashcardToDeck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		d.ew.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}

	var err error
	reqBody := models.AddFlashcardToDeckRequest{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		d.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := d.s.PutAllFlashcards(reqBody.DeckId, []models.Flashcard{reqBody.Flashcard})
	if err != nil {
		d.logger.Errorf("error while putting a card: %v", err)
		d.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = d.s.PutAllUserLeitner([]models.UserLeitner{{
		UserId:      reqBody.UserId,
		FlashcardId: id[0],
		Box:         0,
		CoolDown:    models.CoolDown{State: time.Now()},
	}})
	if err != nil {
		d.logger.Errorf("error while creating leitner: %v", err)
		d.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (d Deck) DeleteFlashcardFromDeck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		d.ew.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}

	reqBody := models.DeleteFlashcardFromDeckRequest{}
	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		d.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := d.s.DeleteFlashcardFromDeck(reqBody.FlashcardId)
	if err != nil {
		d.logger.Errorf("error while deleting flashcard: %v", err)
		d.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (d Deck) NewDeckWithFlashcards(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		d.ew.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}
}

func (d Deck) DeleteDeck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		d.ew.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}
}
