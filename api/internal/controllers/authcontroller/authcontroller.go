package authcontroller

import (
	"api/internal/models/usermodel"
	"net/http"
)

type AuthController struct {
	userModel usermodel.Repository
}

func NewAuthController(u usermodel.Repository) *AuthController {
	return &AuthController{userModel: u}
}

func (c AuthController) SignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func (c AuthController) SignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func (c AuthController) ParseToken(token string) error {
	return nil
}
