package authcontroller

import (
	"api/internal/models/usermodel"
	"net/http"
)

type AuthController struct {
	userModel usermodel.IUser
}

func NewAuthController(u usermodel.IUser) *AuthController {
	return &AuthController{userModel: u}
}

func (c AuthController) SignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func (c AuthController) SignUpShort() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := c.userModel.CreateUserShort()
		if err != nil {
			return
		}
	}
}
func (c AuthController) SignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func (c AuthController) ParseToken(token string) error {
	return nil
}
