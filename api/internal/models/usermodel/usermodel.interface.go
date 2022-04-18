package usermodel

type IUser interface {
	CreateUser(user *User) error
	CreateUserShort() (*User, error)
	BuildUserToken() error
	GetUsers(user *User) ([]*User, error)
}
