package usermodel

import (
	"api/internal/database"
	"api/internal/logger"
	"context"
	"database/sql"
	"time"
)

type User struct {
	ID        uint            `json:"id"`
	Username  *sql.NullString `json:"username,omitempty"`
	Password  string          `json:"-"`
	Email     string          `json:"email"`
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

func (c UserRepository) CreateUser(ctx context.Context, user *User) error {
	return nil
}
func (c UserRepository) GetUsers(ctx context.Context, user *User) ([]*User, error) {
	return nil, nil
}
