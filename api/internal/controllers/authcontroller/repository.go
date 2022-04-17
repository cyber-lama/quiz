package authcontroller

import (
	"net/http"
)

type Repository interface {
	SignUp() http.HandlerFunc
	SignIn() http.HandlerFunc
	ParseToken(token string) error
}
