package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/ogniloud/madr/internal/flashcards/models"
	"github.com/ogniloud/madr/internal/wordmaster"

	"github.com/go-chi/chi/v5"

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
	handlers2 "github.com/ogniloud/madr/internal/social/handlers"
	socialstorage "github.com/ogniloud/madr/internal/social/storage/sql"
	"github.com/ogniloud/madr/internal/usercred"

	"github.com/charmbracelet/log"
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

var initSqlPath = "./db/init.sql"

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

	_initSqlPath := os.Getenv("INIT_SQL_PATH")
	if _initSqlPath != "" {
		initSqlPath = _initSqlPath
	}

	b, err := os.ReadFile(initSqlPath)
	if err == nil {
		_, err = psqlDB.Exec(ctx, string(b))
		if err != nil {
			log.Fatal("init.sql exec error", "error", err)
		}
	} else {
		log.Error("read init.sql fail", "error", err)
	}

	// storage by package definition
	credentials := usercred.New(psqlDB)
	deckStorage := &deckstorage.Storage{Conn: psqlDB}
	socialStorage := &socialstorage.Storage{Conn: psqlDB}

	if err := credentials.ImportGoldenWordsForOld(ctx); err != nil {
		log.Error("Unable to import golden words for old users", "error", err)
	}

	log.Print("Golden words imported")

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

	creq := make(chan *wordmaster.WiktionaryRequest)
	defer close(creq)

	brockers := strings.Split(os.Getenv("BROCKERS"), ",")
	producer := wordmaster.NewProducer(l, brockers, "wiktionary")
	producer.Produce(ctx, creq)

	consumer := wordmaster.NewConsumer(l, brockers, "baked-words", time.Second)
	cresp := consumer.Consume(ctx)

	go func() {
		for msg := range cresp {
			var backsides []models.Backside
			for _, w := range msg.Words {
				for _, s := range w.Senses {
					for _, v := range s.Glosses {
						backsides = append(backsides, models.Backside{
							Value: v,
						})
					}
				}
			}
			if err := deckStorage.AppendBacksides(ctx, models.FlashcardId(msg.Source.CardId), backsides); err != nil {
				l.Error("Unable to append backside to storage", "error", err)
			}
		}
	}()

	deckService := deckserv.NewService(deckStorage, cache.New(), l, creq)
	studyService := study2.NewService(deckService)

	deckEndpoints := deck.New(deckService, ioutil.JSONErrorWriter{Logger: l}, l)
	exerciseEndpoints := study.New(deckService, studyService, ioutil.JSONErrorWriter{Logger: l}, l)

	socialEndpoints := handlers2.New(socialStorage, ioutil.JSONErrorWriter{Logger: l}, l)

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
			r.Get("/card/{id}", deckEndpoints.GetFlashcardById)
			r.Post("/append", deckEndpoints.AppendBacksides)
		})

		// study handler
		r.Route("/study", func(r chi.Router) {
			r.Post("/random", exerciseEndpoints.RandomCard)
			r.Post("/random_deck", exerciseEndpoints.RandomCardDeck)
			r.Post("/rate", exerciseEndpoints.Rate)
			r.Post("/random_matching", exerciseEndpoints.RandomMatching)
			r.Post("/random_matching_deck", exerciseEndpoints.RandomMatchingDeck)
		})

		r.Route("/social", func(r chi.Router) {
			r.Post("/follow", socialEndpoints.Follow)
			r.Post("/unfollow", socialEndpoints.Unfollow)
			r.Post("/followers", socialEndpoints.Followers)
			r.Post("/followings", socialEndpoints.Followings)
			r.Post("/copy", socialEndpoints.DeepCopyDeck)
			r.Get("/search", socialEndpoints.SearchUser)
			r.Post("/feed", socialEndpoints.Feed)
			r.Post("/share", socialEndpoints.ShareWithFollowers)
			r.Post("/is_shared", socialEndpoints.CheckIfSharedWithFollowers)
			r.Post("/groups_shared", socialEndpoints.GetGroupsDeckShared)
		})

		r.Route("/groups", func(r chi.Router) {
			r.Put("/create", socialEndpoints.CreateGroup)
			r.Post("/decks", socialEndpoints.GetDecksByGroupId)
			r.Post("/share", socialEndpoints.ShareGroupDeck)
			r.Post("/delete_deck", socialEndpoints.DeleteGroupDeck)
			r.Get("/search", socialEndpoints.SearchGroupByName)
			r.Post("/groups", socialEndpoints.GetGroupsByUserId)
			r.Post("/created_groups", socialEndpoints.GetCreatedGroupsByUserId)
			r.Put("/change_name", socialEndpoints.ChangeGroupName)
			r.Delete("/delete", socialEndpoints.DeleteGroup)
			r.Delete("/quit", socialEndpoints.QuitGroup)
			r.Post("/participants", socialEndpoints.GetParticipantsByGroupId)
			r.Post("/followers_not_joined", socialEndpoints.GetFollowersNotJoinedGroup)
		})

		r.Route("/invite", func(r chi.Router) {
			r.Post("/send", socialEndpoints.SendInvite)
			r.Post("/accept", socialEndpoints.AcceptInvite)
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
