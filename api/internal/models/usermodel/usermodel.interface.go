package usermodel

import "api/internal/database"

type IUser interface {
	CreateUser(user *User) error
	CreateUserShort() (*User, error)
	BuildUserToken() error
	GetUsers(user *User) ([]*User, error)
}

type IUserToken interface {
	createUserToken(db *database.Database) error
	updateUserToken(db *database.Database) error
	checkUserToken(db *database.Database) (*User, error)
}
