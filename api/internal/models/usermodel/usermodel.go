package usermodel

import (
	"api/internal/database"
	"api/internal/helpers"
	"api/internal/logger"
	"database/sql"
	"time"
)

type User struct {
	ID        uint            `json:"id"`
	Username  *sql.NullString `json:"username,omitempty"`
	Password  *sql.NullString `json:"-"`
	Email     *sql.NullString `json:"email,omitempty"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Token     string          `json:"token"`
}

type UserRepository struct {
	db  *database.Database
	log *logger.Logger
}

func NewUserRepository(db *database.Database, log *logger.Logger) *UserRepository {
	return &UserRepository{
		db:  db,
		log: log,
	}
}

func (r UserRepository) CreateUser(user *User) error {
	return nil
}
func (r UserRepository) CreateUserShort() (*User, error) {
	u := &User{
		CreatedAt: helpers.TimeNow(),
		UpdatedAt: helpers.TimeNow(),
	}
	// Create User
	err := r.db.QueryRow(`
		INSERT INTO users (email, password, created_at, updated_at) VALUES ($1, $2, $3, $4) returning id
	`, u.Email, u.Password, u.CreatedAt, u.UpdatedAt).Scan(&u.ID)

	if err != nil {
		return nil, err
	}
	return u, nil
}
func (r UserRepository) GetUsers(user *User) ([]*User, error) {
	return nil, nil
}
