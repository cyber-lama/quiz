package usermodel

type IUser interface {
	CreateUser(user *User) error
	CreateUserShort() (*User, error)
	GetUsers(user *User) ([]*User, error)
}
