package models

import (
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
