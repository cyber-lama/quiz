package kernel

import (
	"api/internal/configs"
	"api/internal/database"
	"api/internal/logger"
	"github.com/gorilla/mux"
	"log"
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
	// Try to get var from .env
	conf := configs.Init()
	// Try to plug logr
	logr := logger.Init(conf.LogLevel)
	// Try to connect database
	db, err := database.Connect(conf.DBUrl)
	if err != nil {
		//if errors.As(err, &exceptions.ConnectionDBErr{}) {
		//}
		log.Fatal(err)
	}
	return &App{
		configs: conf,
		logger:  logr,
		db:      db,
	}
}
