package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/ogniloud/madr/internal/models"
)

func (e *Endpoints) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		e.ew.Error(w, models.ErrInternalServer.Error(), http.StatusInternalServerError)
		e.logger.Error("Unable to convert user id to integer", "error", err)
		return
	}

	if userID < 1 {
		e.ew.Error(w, "User id cannot be less than 1", http.StatusBadRequest)
		return
	}
}
