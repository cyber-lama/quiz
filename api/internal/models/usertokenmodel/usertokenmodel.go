package usertokenmodel

import "time"

type Token struct {
	ID        uint      `json:"id"`
	UserID    int       `json:"user_id"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
