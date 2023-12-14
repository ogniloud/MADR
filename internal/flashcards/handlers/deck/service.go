package deck

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/ogniloud/madr/internal/flashcards/models"
	"github.com/ogniloud/madr/internal/flashcards/services/deck"
	"github.com/ogniloud/madr/internal/ioutil"

	"github.com/charmbracelet/log"
)

type Endpoints struct {
	s      deck.IService
	ew     ioutil.ErrorWriter
	logger *log.Logger
}

func New(s deck.IService, ew ioutil.ErrorWriter, logger *log.Logger) *Endpoints {
	return &Endpoints{
		s:      s,
		ew:     ew,
		logger: logger,
	}
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

// LoadDecks is a handler for the loading decks Endpoints.
func (d Endpoints) LoadDecks(w http.ResponseWriter, r *http.Request) {
	reqBody := models.LoadDecksRequest{}
	respBody := models.LoadDecksResponse{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		d.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	decksMap, err := d.s.LoadDecks(r.Context(), reqBody.UserId)
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

// GetFlashcardsByDeckId is a handler for the getting cards Endpoints.
func (d Endpoints) GetFlashcardsByDeckId(w http.ResponseWriter, r *http.Request) {
	reqBody := models.GetFlashcardsByDeckIdRequest{}
	respBody := models.GetFlashcardsByDeckIdResponse{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		d.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ids, err := d.s.GetFlashcardsIdByDeckId(r.Context(), reqBody.DeckId)
	if err != nil {
		d.logger.Errorf("error while loading ids of cards: %v", err)
		d.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.Flashcards = make([]models.Flashcard, len(ids))

	// should be replaced with one read transaction
	for i := 0; i < len(ids); i++ {
		card, err := d.s.GetFlashcardById(r.Context(), ids[i])
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

// AddFlashcardToDeck is a handler for the adding a card Endpoints.
func (d Endpoints) AddFlashcardToDeck(w http.ResponseWriter, r *http.Request) {
	reqBody := models.AddFlashcardToDeckRequest{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		d.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := d.s.PutAllFlashcards(r.Context(), reqBody.DeckId, []models.Flashcard{{
		W: reqBody.Word,
		B: reqBody.Backside,
		A: reqBody.Answer,
	}})
	if err != nil {
		d.logger.Errorf("error while putting a card: %v", err)
		d.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = d.s.PutAllUserLeitner(r.Context(), []models.UserLeitner{{
		UserId:      reqBody.UserId,
		FlashcardId: id[0],
		Box:         0,
		CoolDown:    models.CoolDown(time.Now()),
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

// DeleteFlashcardFromDeck is a handler for the deleting a card Endpoints.
func (d Endpoints) DeleteFlashcardFromDeck(w http.ResponseWriter, r *http.Request) {
	reqBody := models.DeleteFlashcardFromDeckRequest{}
	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		d.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := d.s.DeleteFlashcardFromDeck(r.Context(), reqBody.FlashcardId)
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

// NewDeckWithFlashcards is a handler for the creating a new deck Endpoints.
func (d Endpoints) NewDeckWithFlashcards(w http.ResponseWriter, r *http.Request) {
	reqBody := models.NewDeckWithFlashcardsRequest{}
	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		d.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cards := make([]models.Flashcard, len(reqBody.Flashcards))
	for i := 0; i < len(cards); i++ {
		cards[i].W = reqBody.Flashcards[i].Word
		cards[i].B = reqBody.Flashcards[i].Backside
		cards[i].A = reqBody.Flashcards[i].Answer
	}
	log.Print(cards)
	_, err := d.s.NewDeckWithFlashcards(r.Context(),
		reqBody.UserId,
		models.DeckConfig{
			UserId: reqBody.UserId,
			Name:   reqBody.Name,
		},
		cards,
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

// DeleteDeck is a handler for the deleting a deck from user's collection Endpoints.
func (d Endpoints) DeleteDeck(w http.ResponseWriter, r *http.Request) {
	reqBody := models.DeleteDeckRequest{}
	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		d.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := d.s.DeleteDeck(r.Context(), reqBody.UserId, reqBody.DeckId)
	if err != nil {
		d.logger.Errorf("error while deleting a deck: %v", err)
		d.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// swagger:route GET /api/flashcards/card/{id} GetFlashcardById
// Takes a flashcard by id.
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
// + name: id
//   in: query
//   description: FlashcardId.
//   required: true
//   type: int
//
//
// Responses:
// 200: getFlashcardByIdOKResponse
// 400: getFlashcardByIdBadRequestError
// 500: getFlashcardByIdInternalServerError

// GetFlashcardById is a handler that takes a flashcard by id.
func (d Endpoints) GetFlashcardById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		d.ew.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if id < 0 {
		d.ew.Error(w, "id must be >= 0", http.StatusBadRequest)
		return
	}

	respBody := models.GetFlashcardByIdResponse{}
	respBody.Flashcard, err = d.s.GetFlashcardById(r.Context(), id)
	if err != nil {
		d.ew.Error(w, "internal server error", http.StatusInternalServerError)
		d.logger.Error("error on backside", "error", err)
		return
	}

	if err := ioutil.ToJSON(respBody, w); err != nil {
		d.ew.Error(w, "internal server error", http.StatusInternalServerError)
		d.logger.Error("unable to marshal flashcard", "error", err, "body", respBody)
		return
	}
}
