package handlers

import (
	"errors"
	"net/http"

	"github.com/ogniloud/madr/internal/data"
	"github.com/ogniloud/madr/internal/models"
)

// SignUp is a handler for the sign-up endpoint
func (e *Endpoints) SignUp(w http.ResponseWriter, r *http.Request) {
	// We always return JSON from our API
	w.Header().Set("Content-Type", "application/json")

	var request models.SignUpRequest

	// We use the FromJSON function to deserialize the request body
	// because it is faster than using the json.Unmarshal function
	err := models.FromJSON(&request, r.Body)
	if err != nil {
		e.writeGenericError(w, http.StatusBadRequest, "Unable to unmarshal JSON")
		return
	}

	// We create separate models for API request and datalayer request
	// because we don't want to expose the datalayer models to the API
	// users. This is a good practice to follow.
	// And also in case we want to change the datalayer models or the
	// API request models, we can do it without affecting the other.
	user, err := e.data.CreateUser(r.Context(), models.User{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		if errors.Is(err, data.ErrEmailExists) {
			e.writeGenericError(w, http.StatusConflict, data.ErrEmailExists.Error())
			return
		}
		e.logger.Error("Unable to create user", "error", err)
		e.writeGenericError(w, http.StatusInternalServerError, models.ErrInternalServer.Error())
		return
	}

	err = models.ToJSON(models.SignUpResponse{
		ID:    user.ID,
		Email: user.Email,
	}, w)
	if err != nil {
		e.logger.Error("Unable to write JSON response", "error", err)
		e.writeGenericError(w, http.StatusInternalServerError, models.ErrInternalServer.Error())
		return
	}

	// We set the status code to 201 to indicate that the resource is created
	w.WriteHeader(http.StatusCreated)
}

// SignIn is a handler for the sign-in endpoint.
func (e *Endpoints) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request models.SignInRequest

	err := models.FromJSON(&request, r.Body)
	if err != nil {
		e.writeGenericError(w, http.StatusBadRequest, "Unable to unmarshal JSON")
	}

	authToken, err := e.data.SignInUser(r.Context(), request.Email, request.Password)
	if err != nil {
		if errors.Is(err, models.ErrUnauthorized) {
			e.writeGenericError(w, http.StatusUnauthorized, models.ErrUnauthorized.Error())
			return
		}
		e.logger.Error("Unable to get Bearer token", "error", err)
		e.writeGenericError(w, http.StatusInternalServerError, models.ErrInternalServer.Error())
		return
	}

	err = models.ToJSON(models.SignInResponse{
		Authorization: authToken,
	}, w)
	if err != nil {
		// log the error to debug it
		e.logger.Error("Unable to write JSON response", "error", err)
		// write a generic error to the response writer, so we don't expose the actual error
		e.writeGenericError(w, http.StatusInternalServerError, models.ErrInternalServer.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}