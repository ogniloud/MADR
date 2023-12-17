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
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cardId, err := e.ss.GetNextRandom(r.Context(), reqBody.UserId, models.CoolDown(time.Now()))
	if err != nil {
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.Flashcard, err = e.ds.GetFlashcardById(r.Context(), cardId)
	if err != nil {
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := ioutil.ToJSON(&respBody, w); err != nil {
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// swagger:route POST /api/study/random_n RandomNCards
// Returns n random cards from all the decks.
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
//   description: Random N Cards request.
//   required: true
//   type: randomNCardsRequest
//
//
// Responses:
// 200: randomNCardsOkResponse
// 400: randomNCardsBadRequestError
// 500: randomNCardsInternalServerError

// RandomNCards is a handler for getting n random card.
func (e *Endpoints) RandomNCards(w http.ResponseWriter, r *http.Request) {
	reqBody := models.RandomNCardsRequest{}
	respBody := models.RandomNCardsResponse{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cardIds, err := e.ss.GetNextRandomN(r.Context(), reqBody.UserId, models.CoolDown(time.Now()), reqBody.N)
	if err != nil {
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var cards []models.Flashcard
	for _, cardId := range cardIds {
		card, err := e.ds.GetFlashcardById(r.Context(), cardId)
		if err != nil {
			e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
			e.ew.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cards = append(cards, card)
	}
	respBody.Flashcards = cards

	if err := ioutil.ToJSON(&respBody, w); err != nil {
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// swagger:route POST /api/study/random_deck RandomCardDeck
// Returns a random card from the deck.
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
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cardId, err := e.ss.GetNextRandomDeck(r.Context(), reqBody.UserId, reqBody.DeckId, models.CoolDown(time.Now()))
	if err != nil {
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.Flashcard, err = e.ds.GetFlashcardById(r.Context(), cardId)
	if err != nil {
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := ioutil.ToJSON(&respBody, w); err != nil {
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// swagger:route POST /api/study/random_deck_n RandomNCardsDeck
// Returns n random cards from the deck.
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
//   description: Random N Cards Deck request.
//   required: true
//   type: randomNCardsDeckRequest
//
//
// Responses:
// 200: randomNCardsDeckOkResponse
// 400: randomNCardsDeckBadRequestError
// 500: randomNCardsDeckInternalServerError

// RandomNCardsDeck is a handler for getting n random card.
func (e *Endpoints) RandomNCardsDeck(w http.ResponseWriter, r *http.Request) {
	reqBody := models.RandomNCardsDeckRequest{}
	respBody := models.RandomNCardsDeckResponse{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cardIds, err := e.ss.GetNextRandomDeckN(r.Context(), reqBody.UserId, reqBody.DeckId, models.CoolDown(time.Now()), reqBody.N)
	if err != nil {
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var cards []models.Flashcard
	for _, cardId := range cardIds {
		card, err := e.ds.GetFlashcardById(r.Context(), cardId)
		if err != nil {
			e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
			e.ew.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cards = append(cards, card)
	}
	respBody.Flashcards = cards

	if err := ioutil.ToJSON(&respBody, w); err != nil {
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
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
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !isValidMark(reqBody.Mark) {
		e.logger.Errorf("mark is not valid: %v", reqBody.Mark)
		e.ew.Error(w, fmt.Sprintf("mark is not valid: %v", reqBody.Mark), http.StatusBadRequest)
		return
	}

	err := e.ss.Rate(r.Context(), reqBody.UserId, reqBody.FlashcardId, reqBody.Mark)
	if err != nil {
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// swagger:route POST /api/study/random_matching RandomMatching
// Returns a random matching exercise from all the decks.
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
//   description: Random Matching request.
//   required: true
//   type: randomMatchingRequest
//
//
// Responses:
// 200: randomMatchingOkResponse
// 400: randomMatchingBadRequestError
// 500: randomMatchingInternalServerError

// RandomMatching is a handler for getting a random matching exercise.
func (e *Endpoints) RandomMatching(w http.ResponseWriter, r *http.Request) {
	reqBody := models.RandomMatchingRequest{}
	respBody := models.RandomMatchingResponse{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	matching, err := e.ss.MakeMatching(r.Context(), reqBody.UserId, models.CoolDown(time.Now()), reqBody.Size)
	if err != nil {
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.Matching = matching

	if err := ioutil.ToJSON(&respBody, w); err != nil {
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// swagger:route POST /api/study/random_matching_deck RandomMatchingDeck
// Returns a random matching exercise from the deck.
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
//   description: Random Matching Deck request.
//   required: true
//   type: randomMatchingDeckRequest
//
//
// Responses:
// 200: randomMatchingDeckOkResponse
// 400: randomMatchingDeckBadRequestError
// 500: randomMatchingDeckInternalServerError

// RandomMatchingDeck is a handler for getting a random matching exercise.
func (e *Endpoints) RandomMatchingDeck(w http.ResponseWriter, r *http.Request) {
	reqBody := models.RandomMatchingDeckRequest{}
	respBody := models.RandomMatchingDeckResponse{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	matching, err := e.ss.MakeMatchingDeck(r.Context(), reqBody.UserId, reqBody.DeckId, models.CoolDown(time.Now()), reqBody.Size)
	if err != nil {
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.Matching = matching

	if err := ioutil.ToJSON(&respBody, w); err != nil {
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// swagger:route POST /api/study/random_text RandomText
// Returns a random text exercise from all the decks.
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
//   description: Random Text request.
//   required: true
//   type: randomTextRequest
//
//
// Responses:
// 200: randomTextOkResponse
// 400: randomTextBadRequestError
// 500: randomTextInternalServerError

// RandomText is a handler for getting a random text exercise.
func (e *Endpoints) RandomText(w http.ResponseWriter, r *http.Request) {
	reqBody := models.RandomTextRequest{}
	respBody := models.RandomTextResponse{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	text, err := e.ss.MakeText(r.Context(), reqBody.UserId, models.CoolDown(time.Now()), reqBody.Size)
	if err != nil {
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.Text = text

	if err := ioutil.ToJSON(&respBody, w); err != nil {
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// swagger:route POST /api/study/random_text_deck RandomTextDeck
// Returns a random text exercise from the deck.
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
//   description: Random Text Deck request.
//   required: true
//   type: randomTextDeckRequest
//
//
// Responses:
// 200: randomTextDeckOkResponse
// 400: randomTextDeckBadRequestError
// 500: randomTextDeckInternalServerError

// RandomTextDeck is a handler for getting a random text exercise.
func (e *Endpoints) RandomTextDeck(w http.ResponseWriter, r *http.Request) {
	reqBody := models.RandomTextDeckRequest{}
	respBody := models.RandomTextDeckResponse{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	text, err := e.ss.MakeTextDeck(r.Context(), reqBody.UserId, reqBody.DeckId, models.CoolDown(time.Now()), reqBody.Size)
	if err != nil {
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.Text = text

	if err := ioutil.ToJSON(&respBody, w); err != nil {
		e.logger.Errorf("reqBody: %v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func isValidMark(mark models.Mark) bool {
	return mark <= study.Excellent
}
