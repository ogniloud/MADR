package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ogniloud/madr/internal/data"
	"github.com/ogniloud/madr/internal/handlers"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	bindAddress     = ":8080"
	shutdownTimeout = 30 * time.Second
	readTimeout     = 5 * time.Second
	writeTimeout    = 10 * time.Second
	idleTimeout     = 120 * time.Second
)

func main() {
	l := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
	})

	// Set up a router
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Set up a datalayer
	dl := data.NewDatalayer()

	// Set up endpoints
	endpoints := handlers.NewEndpoints(dl, l)

	// Set up routes
	r.Post("/api/signup", endpoints.SignUp)
	r.Post("/api/signin", endpoints.SignIn)

	// create a new server
	s := http.Server{
		Addr:         bindAddress, // configure the bind address
		Handler:      r,           // set the default handler
		ErrorLog:     l.StandardLog(),
		ReadTimeout:  readTimeout,  // max time to read request from the client
		WriteTimeout: writeTimeout, // max time to write response to the client
		IdleTimeout:  idleTimeout,  // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Info("Starting server", "port", bindAddress)

		l.Fatal("Error form server", "error", s.ListenAndServe())
	}()

	// trap interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	sig := <-c
	l.Infof("Got signal: %v", sig)
	l.Infof("Shutting down...")

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	err := s.Shutdown(ctx)
	if err != nil {
		l.Fatal("Error shutting down server", "error", err)
	}
}
