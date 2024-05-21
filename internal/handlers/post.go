package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"unicode"

	"github.com/ogniloud/madr/internal/data"
	"github.com/ogniloud/madr/internal/ioutil"
	"github.com/ogniloud/madr/internal/models"
	"github.com/ogniloud/madr/internal/usercred"
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

	if err := validate(request); err != nil {
		e.ew.Error(w, fmt.Sprintf("sign up: %v", err), http.StatusBadRequest)
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

var _usernameValid = regexp.MustCompile(`^[A-Za-z0-9]+$`)
var _emailValid = regexp.MustCompile(`^[\w-.]+@([\w-]+\.)+[\w-]{2,4}$`)

var _filter = []string{"huy", "pizd", "xuy", "xyu", "pidor",
	"ebl", "gavn", "suka", "manda", "mudak",
	"mydak", "sex", "cekc", "hui", "siski", "jopa",
	"boobs",
}

func verifyUsernameWords(s string) bool {
	for _, f := range _filter {
		if strings.Contains(s, f) {
			return false
		}
	}
	return true
}

func verifyPassword(s string) bool {
	var sevenOrMore, number, upper bool
	letters := 0
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsLetter(c) || c == ' ':
			letters++
		}
	}
	sevenOrMore = len(s) >= 7
	return sevenOrMore && number && upper
}

func validate(req models.SignUpRequest) error {
	if len(req.Username) < 3 || len(req.Username) > 50 {
		return errors.New("username must contain from 3 to 50 symbols")
	}

	if !verifyUsernameWords(req.Username) {
		return errors.New("username contains forbidden symbols")
	}

	if !_usernameValid.MatchString(req.Username) {
		return errors.New("username must contain latin letters or digits")
	}

	if !_emailValid.MatchString(req.Email) {
		return errors.New("email invalid")
	}

	if !verifyPassword(req.Password) {
		return errors.New("a password must be seven or more characters including one uppercase letter," +
			" one special character and alphanumeric characters")
	}

	return nil
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

		if errors.Is(err, usercred.ErrUserNotFound) {
			e.ew.Error(w, usercred.ErrUserNotFound.Error(), http.StatusUnauthorized)
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
