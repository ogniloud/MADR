package handlers

import (
	"errors"
	"net/http"

	"github.com/ogniloud/madr/internal/data"
	"github.com/ogniloud/madr/internal/models"
)

// SignUp is a handler for the sign-up endpoint
func (e *Endpoints) SignUp(w http.ResponseWriter, r *http.Request) {
	var request models.SignUpRequest

	// We use the FromJSON function to deserialize the request body
	// because it is faster than using the json.Unmarshal function
	err := models.FromJSON(&request, r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
		return
	}

	// We create separate models for API request and datalayer request
	// because we don't want to expose the datalayer models to the API
	// users. This is a good practice to follow.
	// And also in case we want to change the datalayer models or the
	// API request models, we can do it without affecting the other.
	user, err := e.data.CreateUser(models.User{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		if errors.Is(err, data.ErrEmailExists) {
			e.writeGenericError(w, http.StatusBadRequest, data.ErrEmailExists.Error())
			return
		} else {
			e.writeGenericError(w, http.StatusInternalServerError, models.ErrInternalServer.Error())
			return
		}
	}

	err = models.ToJSON(models.SignUpResponse{
		ID:    user.ID,
		Email: user.Email,
	}, w)
	if err != nil {
		e.writeGenericError(w, http.StatusInternalServerError, models.ErrInternalServer.Error())
		return
	}
}
