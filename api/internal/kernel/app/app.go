package app

import (
	"api/internal/configs"
	"net/http"
)

type App struct {
	Config *configs.Config
}

func Run() error {
	conf := configs.Run()
	return http.ListenAndServe(conf.Port, nil)
}
