package handlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ogniloud/madr/internal/data"
	"github.com/ogniloud/madr/internal/handlers"
	"github.com/ogniloud/madr/internal/handlers/mocks"
	"github.com/ogniloud/madr/internal/ioutil"
	"github.com/ogniloud/madr/internal/models"

	"github.com/charmbracelet/log"
)

const (
	jsonCredentials = `{"email":"blabla@gmail.com","password":"Aboba228!","username":"boba"}`
)

func TestEndpoints_SignUp(t *testing.T) {
	bodyReader := strings.NewReader(jsonCredentials)

	request, err := http.NewRequest(http.MethodPost, "/api/signup", bodyReader)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
	}

	response := httptest.NewRecorder()

	// database mock
	dl := mocks.NewDatalayer(t)
	dl.EXPECT().CreateUser(request.Context(), models.User{
		Username: "boba",
		Email:    "blabla@gmail.com",
		Password: "Aboba228!",
	}).Return(nil)

	// logger
	l := log.NewWithOptions(io.Discard, log.Options{})

	// endpoints
	s := handlers.New(dl, ioutil.JSONErrorWriter{Logger: l}, l)

	s.SignUp(response, request)

	got := response.Body.String()
	want := ``

	wantStatus := http.StatusCreated

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if response.Code != wantStatus {
		t.Errorf("got %d, want %d", response.Code, wantStatus)
	}
}

func TestEndpoints_SignUpExisting(t *testing.T) {
	firstBodyReader := strings.NewReader(jsonCredentials)

	firstRequest, err := http.NewRequest(http.MethodPost, "/api/signup", firstBodyReader)
	if err != nil {
		t.Errorf("failed to create firstRequest: %v", err)
	}

	firstResponse := httptest.NewRecorder()

	// database mock
	dl := mocks.NewDatalayer(t)
	dl.EXPECT().CreateUser(firstRequest.Context(), models.User{
		Username: "boba",
		Email:    "blabla@gmail.com",
		Password: "Aboba228!",
	}).Return(data.ErrEmailOrUsernameExists)

	// logger
	l := log.NewWithOptions(io.Discard, log.Options{})

	// endpoints
	s := handlers.New(dl, ioutil.JSONErrorWriter{Logger: l}, l)

	// Create a user with the email blabla@gmail
	s.SignUp(firstResponse, firstRequest)

	secondBodyReader := strings.NewReader(jsonCredentials)

	secondRequest, err := http.NewRequest(http.MethodPost, "/api/signup", secondBodyReader)
	if err != nil {
		t.Errorf("failed to create secondRequest: %v", err)
	}

	secondResponse := httptest.NewRecorder()

	// Try to create a user with the same email
	s.SignUp(secondResponse, secondRequest)

	got := secondResponse.Body.String()
	want := `{"message":"user with this email or username already exists"}
`

	wantStatus := http.StatusConflict

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if secondResponse.Code != wantStatus {
		t.Errorf("got %d, want %d", secondResponse.Code, wantStatus)
	}
}

func TestEndpoints_SignUpBadRequest(t *testing.T) {
	jsonBody := `what am I doing with that body oh`
	bodyReader := strings.NewReader(jsonBody)

	request, err := http.NewRequest(http.MethodPost, "/api/signup", bodyReader)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
	}

	response := httptest.NewRecorder()

	// database mocks
	dl := mocks.NewDatalayer(t)

	// logger
	l := log.NewWithOptions(io.Discard, log.Options{})

	// endpoints
	s := handlers.New(dl, ioutil.JSONErrorWriter{Logger: l}, l)

	s.SignUp(response, request)

	got := response.Body.String()
	want := `{"message":"Unable to unmarshal JSON"}
`

	wantStatus := http.StatusBadRequest

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if response.Code != wantStatus {
		t.Errorf("got %d, want %d", response.Code, wantStatus)
	}
}

func TestEndpoints_SignIn(t *testing.T) {
	// create user
	bodyReader := strings.NewReader(jsonCredentials)

	signUpRequest, err := http.NewRequest(http.MethodPost, "/api/signup", bodyReader)
	if err != nil {
		t.Errorf("failed to create signUpRequest: %v", err)
	}

	response := httptest.NewRecorder()

	// database mocks
	dl := mocks.NewDatalayer(t)
	dl.EXPECT().CreateUser(signUpRequest.Context(), models.User{
		Username: "boba",
		Email:    "blabla@gmail.com",
		Password: "Aboba228!",
	}).Return(nil)

	dl.EXPECT().SignInUser(signUpRequest.Context(), "boba", "Aboba228!").
		Return("Bearer blablablaIMATOKENyouAREpoorBASTARD", nil)

	// logger
	l := log.NewWithOptions(io.Discard, log.Options{})

	// endpoints
	s := handlers.New(dl, ioutil.JSONErrorWriter{Logger: l}, l)

	s.SignUp(response, signUpRequest)

	// sign in
	bodyReader = strings.NewReader(jsonCredentials)

	signInRequest, err := http.NewRequest(http.MethodPost, "/api/signin", bodyReader)
	if err != nil {
		t.Errorf("failed to create signInRequest: %v", err)
	}

	response = httptest.NewRecorder()

	s.SignIn(response, signInRequest)

	got := response.Body.String()
	want := `{"authorization":"Bearer blablablaIMATOKENyouAREpoorBASTARD"}
`
	wantStatus := http.StatusOK

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if response.Code != wantStatus {
		t.Errorf("got %d, want %d", response.Code, wantStatus)
	}
}

func TestEndpoints_SignInBadRequest(t *testing.T) {
	jsonBody := `I don't care about you`
	bodyReader := strings.NewReader(jsonBody)

	request, err := http.NewRequest(http.MethodPost, "/api/signin", bodyReader)
	if err != nil {
		t.Errorf("failed to create signInRequest: %v", err)
	}

	response := httptest.NewRecorder()

	// database mock
	dl := mocks.NewDatalayer(t)

	// logger
	l := log.NewWithOptions(io.Discard, log.Options{})

	// endpoints
	s := handlers.New(dl, ioutil.JSONErrorWriter{Logger: l}, l)

	s.SignIn(response, request)

	got := response.Body.String()
	want := `{"message":"Unable to unmarshal JSON"}
`
	wantStatus := http.StatusBadRequest

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if response.Code != wantStatus {
		t.Errorf("got %d, want %d", response.Code, wantStatus)
	}
}

func TestEndpoints_SignInUnauthorized(t *testing.T) {
	bodyReader := strings.NewReader(jsonCredentials)

	signInRequest, err := http.NewRequest(http.MethodPost, "/api/signin", bodyReader)
	if err != nil {
		t.Errorf("failed to create signInRequest: %v", err)
	}

	response := httptest.NewRecorder()

	// database mocks
	dl := mocks.NewDatalayer(t)
	dl.EXPECT().SignInUser(signInRequest.Context(), "boba", "Aboba228!").Return("", models.ErrUnauthorized)

	// logger
	l := log.NewWithOptions(io.Discard, log.Options{})

	// endpoints
	s := handlers.New(dl, ioutil.JSONErrorWriter{Logger: l}, l)

	s.SignIn(response, signInRequest)

	got := response.Body.String()
	want := `{"message":"invalid credentials"}
`
	wantStatus := http.StatusUnauthorized

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if response.Code != wantStatus {
		t.Errorf("got %d, want %d", response.Code, wantStatus)
	}
}
