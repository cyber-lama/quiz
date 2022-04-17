package usermodel

import "context"

type IUser interface {
	CreateUser(ctx context.Context, user *User) error
	GetUsers(ctx context.Context, user *User) ([]*User, error)
}
