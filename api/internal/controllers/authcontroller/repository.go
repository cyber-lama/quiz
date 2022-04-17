package authcontroller

import (
	"api/internal/models/usermodel"
	"context"
)

type Repository interface {
	CreateUser(ctx context.Context) error
	GetUsers(ctx context.Context, user *usermodel.User) ([]*usermodel.User, error)
	DeleteUser(ctx context.Context, user *usermodel.User, id string) error
}
