package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/ogniloud/madr/internal/ioutil"
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

	userInfo, err := e.data.GetUserById(r.Context(), userID)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			e.ew.Error(w, models.ErrUserNotFound.Error(), http.StatusNotFound)
			return
		}
		e.ew.Error(w, models.ErrInternalServer.Error(), http.StatusInternalServerError)
		e.logger.Error("unable to get user info in GetUserInfo", "error", err, "userId", userID)
		return
	}

	err = ioutil.ToJSON(userInfo, w)
	if err != nil {
		e.ew.Error(w, models.ErrInternalServer.Error(), http.StatusInternalServerError)
		e.logger.Error("unable to marshal userInfo in GetUserInfo", "error", err, "userInfo", userInfo)
		return
	}
}
