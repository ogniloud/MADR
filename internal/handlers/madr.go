package handlers

import (
	"github.com/ogniloud/madr/internal/data"

	"github.com/charmbracelet/log"
)

// Endpoints is a struct that defines the handler for the MADR endpoints.
type Endpoints struct {
	data   *data.Datalayer
	logger *log.Logger
}

// NewEndpoints is a constructor for the Endpoints struct.
func NewEndpoints(data *data.Datalayer, logger *log.Logger) *Endpoints {
	return &Endpoints{
		data:   data,
		logger: logger,
	}
}
