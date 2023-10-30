package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ogniloud/madr/internal/data"

	"github.com/charmbracelet/log"
)

func getEndpoints() *Endpoints {
	dl := data.New()
	l := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
	})

	endpoints := New(dl, l)

	return endpoints
}

func TestEndpoints_SignUp(t *testing.T) {
	jsonBody := []byte(`{"email":"blabla@gmail.com","password":"123"}`)
	bodyReader := bytes.NewReader(jsonBody)

	request, err := http.NewRequest(http.MethodPost, "/api/signup", bodyReader)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
	}

	response := httptest.NewRecorder()

	s := getEndpoints()

	s.SignUp(response, request)

	got := response.Body.String()
	want := `{"id":1,"email":"blabla@gmail.com"}
`

	wantStatus := http.StatusCreated

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if response.Code != wantStatus {
		t.Errorf("got %d, want %d", response.Code, wantStatus)
	}
}

func TestEndpoints_SignUpExisting(t *testing.T) {
	jsonBody := []byte(`{"email":"blabla@gmail.com","password":"123"}`)
	firstBodyReader := bytes.NewReader(jsonBody)

	firstRequest, err := http.NewRequest(http.MethodPost, "/api/signup", firstBodyReader)
	if err != nil {
		t.Errorf("failed to create firstRequest: %v", err)
	}

	firstResponse := httptest.NewRecorder()

	s := getEndpoints()

	// Create a user with the email blabla@gmail
	s.SignUp(firstResponse, firstRequest)

	secondBodyReader := bytes.NewReader(jsonBody)

	secondRequest, err := http.NewRequest(http.MethodPost, "/api/signup", secondBodyReader)
	if err != nil {
		t.Errorf("failed to create secondRequest: %v", err)
	}

	secondResponse := httptest.NewRecorder()

	// Try to create a user with the same email
	s.SignUp(secondResponse, secondRequest)

	got := secondResponse.Body.String()
	want := `{"message":"user with this email already exists"}
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
	jsonBody := []byte(`what am I doing with that body oh`)
	bodyReader := bytes.NewReader(jsonBody)

	request, err := http.NewRequest(http.MethodPost, "/api/signup", bodyReader)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
	}

	response := httptest.NewRecorder()

	s := getEndpoints()

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
