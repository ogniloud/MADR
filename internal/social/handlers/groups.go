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
