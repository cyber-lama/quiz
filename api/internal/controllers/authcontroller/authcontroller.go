package authcontroller

import (
	"api/internal/controllers/basecontroller"
	"api/internal/models/usermodel"
	"net/http"
)

type AuthController struct {
	baseCR    basecontroller.IBase
	userModel usermodel.IUser
}

func NewAuthController(u usermodel.IUser, b basecontroller.IBase) *AuthController {
	return &AuthController{userModel: u, baseCR: b}
}

func (c AuthController) SignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func (c AuthController) SignUpShort() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := c.userModel.CreateUserShort()
		if err != nil {
			c.baseCR.Error(w, http.StatusBadRequest, err)
			return
		}
		c.baseCR.Message(w, http.StatusOK, res)
	}
}
func (c AuthController) SignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func (c AuthController) ParseToken(token string) error {
	return nil
}
