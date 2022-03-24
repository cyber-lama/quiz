package app

import (
	"api/internal/configs"
	"api/internal/database"
	"github.com/gorilla/mux"
	"net/http"
)

type App struct {
	configs *configs.Config
	router  *mux.Router
	db      *database.DB
}

func (a App) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	a.router.ServeHTTP(writer, request)
}

func Init() *App {
	conf := configs.Init()
	return &App{
		configs: conf,
		db:      database.Connect(conf.DBUrl),
	}
}
