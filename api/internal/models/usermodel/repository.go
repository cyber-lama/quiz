package usermodel

import "context"

type Repository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUsers(ctx context.Context, user *User) ([]*User, error)
}
