package handlers

import (
	"net/http"

	"github.com/ogniloud/madr/internal/ioutil"
	"github.com/ogniloud/madr/internal/social/models"
)

// ShareGroupDeck /api/group/share
func (e Endpoints) ShareGroupDeck(w http.ResponseWriter, r *http.Request) {
	reqBody := models.ShareGroupDeckRequest{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("json not parsed: %v", err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := e.s.ShareDeckGroup(r.Context(), reqBody.UserId, reqBody.GroupId, reqBody.DeckId)
	if err != nil {
		e.logger.Errorf("reqBody: %+v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (e Endpoints) GetDecksByGroupId(w http.ResponseWriter, r *http.Request) {
	reqBody := models.GetDecksByGroupIdRequest{}
	respBody := models.GetDecksByGroupIdResponse{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("json not parsed: %v", err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	decks, err := e.s.GetDecksByGroupId(r.Context(), reqBody.GroupId)
	if err != nil {
		e.logger.Errorf("reqBody: %+v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.Decks = decks
	if err := ioutil.ToJSON(respBody, w); err != nil {
		e.logger.Errorf("To json error: %v", err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (e Endpoints) DeepCopyDeck(w http.ResponseWriter, r *http.Request) {
	reqBody := models.DeepCopyDeckRequest{}
	respBody := models.DeepCopyDeckResponse{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("json not parsed: %v", err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deckId, err := e.s.DeepCopyDeck(r.Context(), reqBody.CopierId, reqBody.DeckId)
	if err != nil {
		e.logger.Errorf("reqBody: %+v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.DeckId = deckId
	if err := ioutil.ToJSON(respBody, w); err != nil {
		e.logger.Errorf("To json error: %v", err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
