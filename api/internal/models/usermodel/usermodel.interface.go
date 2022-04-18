package usermodel

type IUser interface {
	CreateUser(user *User) error
	CreateUserShort() (*User, error)
	CreateUserToken() error
	GetUsers(user *User) ([]*User, error)
}
