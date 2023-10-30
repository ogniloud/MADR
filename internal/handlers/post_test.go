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

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
