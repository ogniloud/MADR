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

// GetDecksByGroupId /api/group/decks
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

// DeleteGroupDeck /api/groups/delete_deck
func (e Endpoints) DeleteGroupDeck(w http.ResponseWriter, r *http.Request) {
	reqBody := models.DeleteGroupDeckRequest{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("json not parsed: %v", err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := e.s.DeleteDeckFromGroup(r.Context(), reqBody.UserId, reqBody.GroupId, reqBody.DeckId)
	if err != nil {
		e.logger.Errorf("reqBody: %+v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// SearchGroupByName GET /api/groups/search?q=...
func (e Endpoints) SearchGroupByName(w http.ResponseWriter, r *http.Request) {
	respBody := models.SearchGroupByNameResponse{}

	name := r.Form.Get("q")

	groups, err := e.s.GetGroupsByName(r.Context(), name)
	if err != nil {
		e.logger.Errorf("query: %+v, error: %v", r.Form, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.Groups = groups
	if err := ioutil.ToJSON(respBody, w); err != nil {
		e.logger.Errorf("To json error: %v", err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}