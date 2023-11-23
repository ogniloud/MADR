package handlers

import (
	"errors"
	"net/http"

	"github.com/ogniloud/madr/internal/data"
	"github.com/ogniloud/madr/internal/database"
	"github.com/ogniloud/madr/internal/ioutil"
	"github.com/ogniloud/madr/internal/models"
)

// swagger:route POST /api/signup SignUp
// Creates a new user.
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
//	 in: body
//   description: Sign up request.
//   required: true
//   type: signUpRequest
//
//
// Responses:
//  201: signUpCreatedResponse
//  400: signUpBadRequestError
//  409: signUpConflictError
//  500: signUpInternalServerError

// SignUp is a handler for the sign-up endpoint.
func (e *Endpoints) SignUp(w http.ResponseWriter, r *http.Request) {
	// We always return JSON from our API
	w.Header().Set("Content-Type", "application/json")

	var request models.SignUpRequest

	// We use the FromJSON function to deserialize the request body
	// because it is faster than using the json.Unmarshal function
	err := ioutil.FromJSON(&request, r.Body)
	if err != nil {
		e.ew.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
		return
	}

	// We create separate models for API request and datalayer request
	// because we don't want to expose the datalayer models to the API
	// users. This is a good practice to follow.
	// And also in case we want to change the datalayer models or the
	// API request models, we can do it without affecting the other.
	err = e.data.CreateUser(r.Context(), models.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		if errors.Is(err, data.ErrEmailOrUsernameExists) {
			e.ew.Error(w, data.ErrEmailOrUsernameExists.Error(), http.StatusConflict)
			return
		}

		e.logger.Error("Unable to create user", "error", err)
		e.ew.Error(w, models.ErrInternalServer.Error(), http.StatusInternalServerError)

		return
	}

	// We set the status code to 201 to indicate that the resource is created
	w.WriteHeader(http.StatusCreated)
}

// swagger:route POST /api/signin SignIn
// Signs in a user.
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
//   description: Sign in request.
//	 required: true
//	 type: signInRequest
//
//
// Responses:
// 200: signInOkResponse
// 400: signInBadRequestError
// 401: signInUnauthorizedError
// 500: signInInternalServerError

// SignIn is a handler for the sign-in endpoint.
func (e *Endpoints) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request models.SignInRequest

	err := ioutil.FromJSON(&request, r.Body)
	if err != nil {
		e.ew.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
		return
	}

	authToken, err := e.data.SignInUser(r.Context(), request.Username, request.Password)
	if err != nil {
		if errors.Is(err, models.ErrUnauthorized) {
			e.ew.Error(w, models.ErrUnauthorized.Error(), http.StatusUnauthorized)
			return
		}

		if errors.Is(err, database.ErrUserNotFound) {
			e.ew.Error(w, database.ErrUserNotFound.Error(), http.StatusUnauthorized)
			return
		}

		e.logger.Error("Unable to get Bearer token", "error", err)
		e.ew.Error(w, models.ErrInternalServer.Error(), http.StatusInternalServerError)

		return
	}

	err = ioutil.ToJSON(models.SignInResponse{
		Authorization: authToken,
	}, w)
	if err != nil {
		// log the error to debug it
		e.logger.Error("Unable to write JSON response", "error", err)
		// write a generic error to the response writer, so we don't expose the actual error
		e.ew.Error(w, models.ErrInternalServer.Error(), http.StatusInternalServerError)

		return
	}
}
