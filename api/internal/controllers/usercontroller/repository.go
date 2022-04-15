package usercontroller

import (
	"api/internal/models"
	"context"
)

type Repository interface {
	CreateUser(ctx context.Context) error
	DeleteUser(ctx context.Context, user *models.User, id string) error
}
