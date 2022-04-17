package authcontroller

import (
	"net/http"
)

type IAuth interface {
	SignUp() http.HandlerFunc
	SignIn() http.HandlerFunc
	ParseToken(token string) error
}
