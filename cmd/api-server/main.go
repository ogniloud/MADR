package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	_ "github.com/lib/pq"

	"github.com/go-chi/cors"

	"github.com/ogniloud/madr/internal/data"
	"github.com/ogniloud/madr/internal/db"
	"github.com/ogniloud/madr/internal/flashcards/cache"
	"github.com/ogniloud/madr/internal/flashcards/handlers/deck"
	"github.com/ogniloud/madr/internal/flashcards/handlers/study"
	deckserv "github.com/ogniloud/madr/internal/flashcards/services/deck"
	study2 "github.com/ogniloud/madr/internal/flashcards/services/study"
	deckstorage "github.com/ogniloud/madr/internal/flashcards/storage/deck"
	"github.com/ogniloud/madr/internal/handlers"
	"github.com/ogniloud/madr/internal/ioutil"
	"github.com/ogniloud/madr/internal/usercred"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	openapi "github.com/go-openapi/runtime/middleware"
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

	// Set up a routerDeck
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(
		cors.Options{
			AllowedOrigins: []string{"https://*", "http://*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		}))

	// Set up a context
	ctx := context.Background()

	// Get the connection string from the environment
	connStr := os.Getenv("DB_CONN_STR")
	if connStr == "" {
		l.Fatal("DB_CONN_STR env var is not set")
	}

	// wait until database is up
	time.Sleep(5 * time.Second)

	// Set up a database connection
	psqlDB, err := db.NewPSQLDatabase(ctx, connStr, l)
	if err != nil {
		log.Fatal(err)
	}

	// set up context so that we can ping the database and don't wait forever
	cancelCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// check whether connection is established
	err = psqlDB.Ping(cancelCtx)
	if err != nil {
		log.Fatal("Unable to ping database in main", "error", err)
	}

	// storage by package definition
	credentials := usercred.New(psqlDB)
	deckStorage := &deckstorage.Storage{Conn: psqlDB}

	// get salt length from env
	saltLengthString := os.Getenv("SALT_LENGTH")
	if saltLengthString == "" {
		l.Fatal("SALT_LENGTH env var is not set")
	}

	// convert to int
	saltLength, err := strconv.Atoi(saltLengthString)
	if err != nil {
		l.Fatal("Unable to convert salt length to int", "error", err)
	}

	// get token expiration time from env
	tokenExpirationTime := os.Getenv("TOKEN_EXPIRATION_TIME")
	if tokenExpirationTime == "" {
		l.Fatal("TOKEN_EXPIRATION_TIME env var is not set")
	}

	// convert to int
	tokenExpirationTimeInt, err := strconv.Atoi(tokenExpirationTime)
	if err != nil {
		l.Fatal("Unable to convert token expiration time to int", "error", err)
	}

	// convert to time.Duration
	duration := time.Duration(tokenExpirationTimeInt) * time.Hour

	// get sign key from env
	tokenSignKey := os.Getenv("TOKEN_SECRET")
	if tokenSignKey == "" {
		l.Fatal("TOKEN_SECRET env var is not set")
	}

	// Set up a datalayer
	dl := data.New(credentials, saltLength, duration, []byte(tokenSignKey))

	// Set up error writer
	ew := ioutil.JSONErrorWriter{Logger: l}

	// Set up endpoints
	endpoints := handlers.New(dl, ew, l)

	// handlers for documentation
	dh := openapi.Redoc(openapi.RedocOpts{
		BasePath: "/api/",
		SpecURL:  "/api/swagger.yaml",
	}, nil)

	deckService := deckserv.NewService(deckStorage, cache.New(), l)
	studyService := study2.NewService(deckService)

	deckEndpoints := deck.New(deckService, ioutil.JSONErrorWriter{Logger: l}, l)
	exerciseEndpoints := study.New(deckService, studyService, ioutil.JSONErrorWriter{Logger: l}, l)

	// Set up routes
	r.Route("/api", func(r chi.Router) {
		r.Post("/signup", endpoints.SignUp)
		r.Options("/signup", endpoints.SignUp)
		r.Post("/signin", endpoints.SignIn)
		r.Get("/swagger.yaml", http.StripPrefix("/api/", http.FileServer(http.Dir("./"))).ServeHTTP)
		r.Get("/docs", dh.ServeHTTP)
		r.Get("/user/{id}", endpoints.GetUserInfo)

		// deck service handler
		r.Route("/flashcards", func(r chi.Router) {
			r.Put("/add_card", deckEndpoints.AddFlashcardToDeck)
			r.Delete("/delete_deck", deckEndpoints.DeleteDeck)
			r.Delete("/delete_card", deckEndpoints.DeleteFlashcardFromDeck)
			r.Post("/cards", deckEndpoints.GetFlashcardsByDeckId)
			r.Post("/load", deckEndpoints.LoadDecks)
			r.Put("/new_deck", deckEndpoints.NewDeckWithFlashcards)
		})

		// study handler
		r.Route("/study", func(r chi.Router) {
			r.Post("/random", exerciseEndpoints.RandomCard)
			r.Post("/random_deck", exerciseEndpoints.RandomCardDeck)
			r.Post("/rate", exerciseEndpoints.Rate)
		})
	})

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
	cancelCtx, cancel = context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	err = s.Shutdown(cancelCtx)
	if err != nil {
		l.Fatal("Error shutting down server", "error", err)
	}
}
