package usermodel

import (
	"api/internal/database"
	"time"
)

type Token struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t *Token) createUserToken(db *database.Database) error {
	err := db.QueryRow(`
		INSERT INTO user_token (user_id, token, created_at, updated_at) VALUES ($1, $2, $3, $4) returning id
	`, t.UserID, t.Token, t.CreatedAt, t.UpdatedAt).Scan(&t.ID)

	if err != nil {
		return err
	}
	return nil
}

func (t *Token) updateUserToken(db *database.Database) error {
	return nil
}
func (t *Token) checkUserToken(db *database.Database) (*User, error) {
	return nil, nil
}
