package handlers

import (
	"context"

	"github.com/ogniloud/madr/internal/ioutil"
	"github.com/ogniloud/madr/internal/models"

	"github.com/charmbracelet/log"
)

// Datalayer is an interface that defines the methods for the datalayer.
//
//go:generate mockery --name Datalayer --output=./ --filename=mocks/datalayer.go --with-expecter
type Datalayer interface {
	CreateUser(ctx context.Context, user models.User) error
	SignInUser(ctx context.Context, username, password string) (string, error)
}

// Endpoints is a struct that defines the handler for the MADR endpoints.
type Endpoints struct {
	data   Datalayer
	ew     ioutil.ErrorWriter
	logger *log.Logger
}

// New is a constructor for the Endpoints struct.
func New(data Datalayer, logger *log.Logger) *Endpoints {
	return &Endpoints{
		data:   data,
		ew:     ioutil.JSONErrorWriter{Logger: logger},
		logger: logger,
	}
}
