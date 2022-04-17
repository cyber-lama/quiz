package server

import (
	"api/internal/config"
	"api/internal/controllers/authcontroller"
	"api/internal/database"
	"api/internal/logger"
	"api/internal/models/usermodel"
	"net/http"
)

type App struct {
	httpServer *http.Server
	log        *logger.Logger

	authCR authcontroller.Repository
}

func NewApp(conf *config.Config) (*App, error) {
	log := logger.Init(conf.LogLevel)
	db, err := database.Connect(conf.DBUrl)
	if err != nil {
		return nil, err
	}

	// init models
	u := usermodel.NewUserRepository(db, log)
	// init controllers
	app := &App{
		log: log,

		authCR: authcontroller.NewAuthController(u),
	}
	return app, nil
}
