package handlers

import (
	"github.com/ogniloud/madr/internal/data"
	"github.com/ogniloud/madr/internal/ioutil"

	"github.com/charmbracelet/log"
)

// Endpoints is a struct that defines the handler for the MADR endpoints.
type Endpoints struct {
	data   *data.Datalayer
	ew     ioutil.ErrorWriter
	logger *log.Logger
}

// New is a constructor for the Endpoints struct.
func New(data *data.Datalayer, logger *log.Logger) *Endpoints {
	return &Endpoints{
		data:   data,
		ew:     ioutil.JSONErrorWriter{},
		logger: logger,
	}
}
