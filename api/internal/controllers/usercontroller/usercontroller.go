package usercontroller

import (
	"api/internal/controllers/basecontroller"
	"api/internal/database"
	"api/internal/logger"
)

type UserController struct {
	base   *basecontroller.Repository
	db     *database.Database
	logger *logger.Logger
}
