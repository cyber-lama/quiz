package authcontroller

import (
	"api/internal/models/usermodel"
	"context"
)

type AuthController struct {
	userModel usermodel.Repository
}

func NewAuthController(u usermodel.Repository) *AuthController {
	return &AuthController{userModel: u}
}

func (c AuthController) CreateUser(ctx context.Context) error {
	return nil
}
func (c AuthController) GetUsers(ctx context.Context, user *usermodel.User) ([]*usermodel.User, error) {
	return nil, nil
}
func (c AuthController) DeleteUser(ctx context.Context, user *usermodel.User, id string) error {
	return nil
}
