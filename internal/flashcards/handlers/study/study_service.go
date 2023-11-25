package study

import (
	"fmt"
	"net/http"
	"time"

	"github.com/charmbracelet/log"

	"github.com/ogniloud/madr/internal/flashcards/models"
	"github.com/ogniloud/madr/internal/flashcards/services/deck"
	"github.com/ogniloud/madr/internal/flashcards/services/study"
	"github.com/ogniloud/madr/internal/ioutil"
)

type Endpoints struct {
	ds     deck.IService
	ss     study.IStudyService
	ew     ioutil.ErrorWriter
	logger *log.Logger
}

func New(ds deck.IService, ss study.IStudyService, ew ioutil.ErrorWriter, logger *log.Logger) *Endpoints {
	return &Endpoints{
		ds:     ds,
		ss:     ss,
		ew:     ew,
		logger: logger,
	}
}

// swagger:route POST /api/study/random RandomCard
// Returns a random card from all the decks.
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
//   description: Random Card request.
//   required: true
//   type: randomCardRequest
//
//
// Responses:
// 200: randomCardOkResponse
// 400: randomCardBadRequestError
// 500: randomCardInternalServerError

// RandomCard is a handler for getting a random card.
func (e *Endpoints) RandomCard(w http.ResponseWriter, r *http.Request) {
	reqBody := models.RandomCardRequest{}
	respBody := models.RandomCardResponse{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cardId, err := e.ss.GetNextRandom(r.Context(), reqBody.UserId, models.CoolDown(time.Now()))
	if err != nil {
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.Flashcard, err = e.ds.GetFlashcardById(r.Context(), cardId)
	if err != nil {
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := ioutil.ToJSON(&respBody, w); err != nil {
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// swagger:route POST /api/study/random_deck RandomCardDeck
// Returns a random card from the decks.
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
//   description: Random Card Deck request.
//   required: true
//   type: randomCardDeckRequest
//
//
// Responses:
// 200: randomCardDeckOkResponse
// 400: randomCardDeckBadRequestError
// 500: randomCardDeckInternalServerError

// RandomCardDeck is a handler for getting a random card from the deck.
func (e *Endpoints) RandomCardDeck(w http.ResponseWriter, r *http.Request) {
	reqBody := models.RandomCardDeckRequest{}
	respBody := models.RandomCardDeckResponse{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cardId, err := e.ss.GetNextRandomDeck(r.Context(), reqBody.UserId, reqBody.DeckId, models.CoolDown(time.Now()))
	if err != nil {
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.Flashcard, err = e.ds.GetFlashcardById(r.Context(), cardId)
	if err != nil {
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := ioutil.ToJSON(&respBody, w); err != nil {
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// swagger:route POST /api/study/rate Rate
// Rates the card and puts a new temperature for the card in a leitner's system.
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
//   description: Rate request.
//   required: true
//   type: rateRequest
//
//
// Responses:
// 204: rateNoContentResponse
// 400: rateBadRequestError
// 500: rateInternalServerError

// Rate is a handler for rating a card in Leitner's system.
func (e *Endpoints) Rate(w http.ResponseWriter, r *http.Request) {
	reqBody := models.RateRequest{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !isValidMark(reqBody.Mark) {
		e.ew.Error(w, fmt.Sprintf("mark is not valid: %v", reqBody.Mark), http.StatusBadRequest)
		return
	}

	err := e.ss.Rate(r.Context(), reqBody.UserId, reqBody.FlashcardId, reqBody.Mark)
	if err != nil {
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func isValidMark(mark models.Mark) bool {
	if mark <= study.Excellent {
		return true
	}

	return false
}
