package handlers

import (
	"net/http"

	"github.com/ogniloud/madr/internal/models"
)

// writeGenericError is a helper function to write a generic error
// to the response writer with the given status code and message.
func (e *Endpoints) writeGenericError(w http.ResponseWriter, status int, message string) {
	// write http status code
	w.WriteHeader(status)

	// write json response
	err := models.ToJSON(models.GenericError{
		Message: message,
	}, w)
	if err != nil {
		e.logger.Error("Unable to write JSON response", "error", err)
	}
}
