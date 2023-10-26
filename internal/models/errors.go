package models

import (
	"errors"
)

// ErrInternalServer is a generic error message returned by a server
// in case of an internal server error when we don't want to expose
// the real error to the client. All the internal server errors
// should be logged and fixed. Don't use this error if it's
// something that can be fixed by the client.
var (
	ErrInternalServer = errors.New("don't worry, we are working on it")
	ErrUnauthorized   = errors.New("invalid credentials")
)

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}
