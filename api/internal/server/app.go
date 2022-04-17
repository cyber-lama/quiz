package server

import (
	"api/internal/config"
	"api/internal/controllers/authcontroller"
	"api/internal/database"
	"api/internal/logger"
	"api/internal/models/usermodel"
	"context"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	httpServer *http.Server
	log        *logger.Logger
	router     *mux.Router

	authCR authcontroller.IAuth
}

func NewApp(conf *config.Config) (*App, error) {
	// init logger, default level "fatal"
	log := logger.Init(conf.LogLevel)
	// init db
	db, err := database.Connect(conf.DBUrl)
	if err != nil {
		return nil, err
	}
	// init router
	r := mux.NewRouter()
	// init models
	u := usermodel.NewUserRepository(db, log)
	// init controllers
	app := &App{
		log:    log,
		router: r,

		authCR: authcontroller.NewAuthController(u),
	}
	return app, nil
}

func (a *App) Run(port string) error {
	a.configureRouter()

	httpServer := &http.Server{
		Addr:           port,
		Handler:        a.router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			a.log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()
	// on ctrl+c stop server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
	err := httpServer.Shutdown(ctx)
	if err != nil {
		a.log.Errorf("Failed shutdown serve: %+v", err)
	}
	return nil
}

// configure the router, add middleware
const (
	ctxKeyRequestID = iota
)

func (a *App) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

func (a *App) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logger.Init("trace")
		log.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxKeyRequestID),
		})

		var payload interface{}
		if r.Method == http.MethodPost {
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Error(err)
			}
			payload = b
		} else if r.Method == http.MethodGet {
			payload = r.URL.Query().Encode()
		}

		log.Infof("--------------------------------")
		log.Infof("started %s %s, with payload: %s",
			r.Method, r.RequestURI, payload)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		var level logrus.Level
		switch {
		case rw.code >= 500:
			level = logrus.ErrorLevel
		case rw.code >= 400:
			level = logrus.WarnLevel
		default:
			level = logrus.InfoLevel
		}
		log.Logf(
			level,
			"completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start),
		)
		log.Infof("--------------------------------")
	})
}

func (a *App) configureRouter() {
	a.router.Use(a.setRequestID)
	a.router.Use(a.logRequest)
	a.router.HandleFunc("/registration", a.authCR.SignUp()).Methods("POST")
}
