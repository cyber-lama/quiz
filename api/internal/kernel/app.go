package kernel

import (
	"api/internal/configs"
	"api/internal/database"
	"api/internal/logger"
	"api/internal/router"
	"log"
	"net/http"
)

type App struct {
	configs *configs.Config
	router  *router.Router
	db      *database.Database
	logger  *logger.Logger
}

func (a App) Run() {
	err := http.ListenAndServe(a.configs.Port, a.router)
	if err != nil {
		a.logger.Fatal(err)
	}
	a.logger.Infof("server start on %s port", a.configs.Port)
}

func (a App) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	a.router.ServeHTTP(writer, request)
}

func Init() *App {
	// Get var from .env
	configsInit := configs.Init()
	// Plug logger
	loggerInit := logger.Init(configsInit.LogLevel)
	// Try to connect database
	dbInit, err := database.Connect(configsInit.DBUrl)
	if err != nil {
		//if errors.As(err, &exceptions.ConnectionDBErr{}) {
		//}
		log.Fatal(err)
	}
	// Init router with sum settings
	routerInit := router.Init()

	return &App{
		configs: configsInit,
		router:  routerInit,
		logger:  loggerInit,
		db:      dbInit,
	}
}
