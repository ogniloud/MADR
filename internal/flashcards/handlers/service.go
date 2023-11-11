package handlers

import (
	"net/http"
	"time"

	"github.com/ogniloud/madr/internal/flashcards/models"
	"github.com/ogniloud/madr/internal/flashcards/services/deck"
	"github.com/ogniloud/madr/internal/ioutil"

	"github.com/charmbracelet/log"
)

type Endpoint struct {
	s      deck.IService
	ew     ioutil.ErrorWriter
	logger *log.Logger
}

// swagger:route POST /api/flashcards/load LoadDecks
// Loads decks that the user has.
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// Schemes: http
//
// Parameters:
// + name: request
//   in: body
//   description: Load decks request.
//   required: true
//   type: loadDecksRequest
//
//
// Responses:
// 200: loadDecksOkResponse
// 400: loadDecksBadRequestError
// 500: loadDecksInternalServerError

// LoadDecks is a handler for the loading decks Endpoint.
func (d Endpoint) LoadDecks(w http.ResponseWriter, r *http.Request) {
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

// swagger:route POST /api/flashcards/cards CardsByDeckId
// Returns flashcards containing in the deck.
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// Schemes: http
//
// Parameters:
// + name: request
//   in: body
//   description: Get flashcards by deck id request.
//   required: true
//   type: getFlashcardsByDeckIdRequest
//
//
// Responses:
// 200: getFlashcardsByDeckIdOkResponse
// 400: getFlashcardsByDeckIdBadRequestError
// 500: getFlashcardsByDeckIdInternalServerError

// GetFlashcardsByDeckId is a handler for the getting cards Endpoint.
func (d Endpoint) GetFlashcardsByDeckId(w http.ResponseWriter, r *http.Request) {
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

// swagger:route PUT /api/flashcards/add_card AddCard
// Puts a card to the deck.
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// Schemes: http
//
// Parameters:
// + name: request
//   in: body
//   description: Add flashcard to the deck request.
//   required: true
//   type: addFlashcardToDeckRequest
//
//
// Responses:
// 201: addFlashcardToDeckCreatedResponse
// 400: addFlashcardToDeckBadRequestError
// 500: addFlashcardToDeckInternalServerError

// AddFlashcardToDeck is a handler for the adding a card Endpoint.
func (d Endpoint) AddFlashcardToDeck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		d.ew.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}

	reqBody := models.AddFlashcardToDeckRequest{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		d.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := d.s.PutAllFlashcards(reqBody.DeckId, []models.Flashcard{{
		W:      reqBody.Word,
		B:      reqBody.Backside,
		DeckId: reqBody.DeckId,
	}})
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

	w.WriteHeader(http.StatusCreated)
}

// swagger:route DELETE /api/flashcards/delete_card DeleteCard
// Deletes a card from the deck.
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// Schemes: http
//
// Parameters:
// + name: request
//   in: body
//   description: Delete card from deck request.
//   required: true
//   type: deleteFlashcardFromDeckRequest
//
//
// Responses:
// 204: deleteFlashcardFromDeckNoContentResponse
// 400: deleteFlashcardFromDeckBadRequestError
// 500: deleteFlashcardFromDeckInternalServerError

// DeleteFlashcardFromDeck is a handler for the deleting a card Endpoint.
func (d Endpoint) DeleteFlashcardFromDeck(w http.ResponseWriter, r *http.Request) {
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

// swagger:route PUT /api/flashcards/new_deck NewDeck
// Creates a new deck with flashcards.
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// Schemes: http
//
// Parameters:
// + name: request
//   in: body
//   description: New deck with flashcards request.
//   required: true
//   type: newDeckWithFlashcardsRequest
//
//
// Responses:
// 201: newDeckWithFlashcardsCreatedResponse
// 400: newDeckWithFlashcardsBadRequestError
// 500: newDeckWithFlashcardsInternalServerError

// NewDeckWithFlashcards is a handler for the creating a new deck Endpoint.
func (d Endpoint) NewDeckWithFlashcards(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		d.ew.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}

	reqBody := models.NewDeckWithFlashcardsRequest{}
	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		d.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := d.s.NewDeckWithFlashcards(
		reqBody.UserId,
		models.DeckConfig{
			UserId: reqBody.UserId,
			Name:   reqBody.Name,
		},
		reqBody.Flashcards,
	)
	if err != nil {
		d.logger.Errorf("error while creating a deck: %v", err)
		d.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// swagger:route DELETE /api/flashcards/delete_deck DeleteDeck
// Deletes a deck from user's collection.
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// Schemes: http
//
// Parameters:
// + name: request
//   in: body
//   description: Delete deck request.
//   required: true
//   type: deleteDeckRequest
//
//
// Responses:
// 204: deleteDeckNoContentResponse
// 400: deleteDeckBadRequestError
// 500: deleteDeckInternalServerError

// DeleteDeck is a handler for the deleting a deck from user's collection Endpoint.
func (d Endpoint) DeleteDeck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		d.ew.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}

	reqBody := models.DeleteDeckRequest{}
	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		d.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := d.s.DeleteDeck(reqBody.UserId, reqBody.DeckId)
	if err != nil {
		d.logger.Errorf("error while deleting a deck: %v", err)
		d.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
