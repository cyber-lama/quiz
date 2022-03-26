package app

import (
	"api/internal/configs"
	"api/internal/database"
	"api/internal/logger"
	"github.com/gorilla/mux"
	"net/http"
)

type App struct {
	configs *configs.Config
	router  *mux.Router
	db      *database.DB
	logger  *logger.Logger
}

func (a App) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	a.router.ServeHTTP(writer, request)
}

func Init() *App {
	conf := configs.Init()
	log := logger.Init(conf.LogLevel)
	return &App{
		configs: conf,
		logger:  log,
		db:      database.Connect(conf.DBUrl),
	}
}
