package usermodel

import (
	"app/internal/database"
	"app/internal/helpers"
	"app/internal/logger"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type nullString struct {
	*sql.NullString
}

type User struct {
	ID        uint        `json:"id"`
	Username  *nullString `json:"username,omitempty"`
	Password  *nullString `json:"-"`
	Email     *nullString `json:"email,omitempty"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Token     string      `json:"token"`
}

type UserRepository struct {
	user *User
	db   *database.Database
	log  *logger.Logger
}

func NewUserRepository(db *database.Database, log *logger.Logger) *UserRepository {
	return &UserRepository{
		db:  db,
		log: log,
	}
}

func (u UserRepository) CreateUser(user *User) error {
	return nil
}
func (u UserRepository) CreateUserShort() (*User, error) {
	u.user = &User{
		CreatedAt: helpers.TimeNow(),
		UpdatedAt: helpers.TimeNow(),
	}
	// Create User
	err := u.db.QueryRow(`
		INSERT INTO users (email, password, created_at, updated_at) VALUES ($1, $2, $3, $4) returning id
	`, u.user.Email, u.user.Password, u.user.CreatedAt, u.user.UpdatedAt).Scan(&u.user.ID)
	if err != nil {
		u.log.Errorf("Ошибка sql запроса, при создании пользователя %s", err)
		return nil, ErrCreateUser
	}

	err = u.BuildUserToken()
	if err != nil {
		u.log.Errorf("Ошибка sql запроса, при создании токена %s", err)
		return nil, ErrCreateUser
	}
	return u.user, nil
}
func (u UserRepository) GetUsers(user *User) ([]*User, error) {
	return nil, nil
}

func (u UserRepository) BuildUserToken() error {
	token := fmt.Sprintf("%d|%s", u.user.ID, uuid.New().String())
	t := Token{
		UserID:    u.user.ID,
		Token:     token,
		CreatedAt: helpers.TimeNow(),
		UpdatedAt: helpers.TimeNow(),
	}
	err := t.createUserToken(u.db)
	if err != nil {
		return err
	}
	u.user.Token = token
	return nil
}
