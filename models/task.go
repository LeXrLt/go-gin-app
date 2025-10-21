package models

import "time"

type Task struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Prompt	  string    `json:"prompt"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ExpiresAt time.Time `json:"expires_at"`
}