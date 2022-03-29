package kernel

import (
	"api/internal/configs"
	"api/internal/database"
	"api/internal/logger"
	"api/internal/router"
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	configs *configs.Config
	router  *router.Router
	db      *database.Database
	logger  *logger.Logger
}

func (a *App) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	a.router.ServeHTTP(writer, request)
}

func (a *App) Run() {
	httpServer := &http.Server{
		Addr:           a.configs.Port,
		Handler:        a.router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		defer a.db.CloseDBConnect()
		if err := httpServer.ListenAndServe(); err != nil {
			a.logger.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	// on ctrl+c stop server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
	defer a.db.CloseDBConnect()
	err := httpServer.Shutdown(ctx)
	if err != nil {
		a.logger.Errorf("Failed shutdown serve: %+v", err)
	}
}

// Init a function that initializes
//all the important components of the application:
//routing, configuration from the env file,
//routing and processing http requests,
//connecting to the database
func Init() (*App, error) {
	// Get var from .env
	c := configs.Init()
	// Plug logger
	l := logger.Init(c.LogLevel)
	// Try to connect database
	db, err := database.Connect(c.DBUrl)
	if err != nil {
		return nil, err
	}
	// Init router with sum settings
	r := router.Init()

	return &App{
		configs: c,
		router:  r,
		logger:  l,
		db:      db,
	}, nil
}
