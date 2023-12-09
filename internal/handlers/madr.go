package handlers

import (
	"context"
	"net/http"

	"github.com/ogniloud/madr/internal/models"
)

// Datalayer is an interface that defines the methods for the datalayer.
//
//go:generate mockery --name Datalayer --output=./ --filename=mocks/datalayer.go --with-expecter
type Datalayer interface {
	CreateUser(ctx context.Context, user models.User) error
	SignInUser(ctx context.Context, username, password string) (string, error)
	GetUserById(ctx context.Context, id int) (models.UserInfo, error)
}

// ErrorWriter is an interface that defines the methods for the error writer.
type ErrorWriter interface {
	Error(w http.ResponseWriter, msg string, status int)
}

// Logger is an interface that defines the methods for the logger.
//
//go:generate mockery --name Logger --output=./ --filename=mocks/logger.go --with-expecter
type Logger interface {
	Error(msg interface{}, keyvals ...interface{})
}

// Endpoints is a struct that defines the handler for the MADR endpoints.
type Endpoints struct {
	data   Datalayer
	ew     ErrorWriter
	logger Logger
}

// New is a constructor for the Endpoints struct.
func New(data Datalayer, ew ErrorWriter, logger Logger) *Endpoints {
	return &Endpoints{
		data:   data,
		ew:     ew,
		logger: logger,
	}
}
