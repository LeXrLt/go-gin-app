package models

import "time"

type SubTask struct {
	ID        int64     `json:"id"`
	TaskID    int64     `json:"task_id"`
	Prompt    string    `json:"prompt"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}