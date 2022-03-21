package app

import (
	"api/internal/configs"
	"github.com/gorilla/mux"
	"net/http"
)

type App struct {
	configs *configs.Config
	router  *mux.Router
}

func (a App) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	a.router.ServeHTTP(writer, request)
}

func Run() error {
	app := &App{
		configs: configs.Run(),
	}
	return http.ListenAndServe(app.configs.Port, app)
}
