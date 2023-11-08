package ioutil

import (
	"net/http"
)

// ErrorWriter is an interface for writing a generic errors
// like http.Error func.
type ErrorWriter interface {
	// Error writes msg to http.ResponseWriter and the status code.
	Error(w http.ResponseWriter, msg string, status int)
}

// GenericError is a generic error message returned by a server.
type GenericError struct {
	// The error message.
	//
	// example: Very useful error message
	Message string `json:"message"`
}
