package models

import "time"

type DistillationData struct {
	ID               int       `json:"id"`
	UserID           int       `json:"user_id"`
	Prompt           string    `json:"prompt"`
	InferenceProcess string    `json:"inference_process"`
	ModelOutput      string    `json:"model_output"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}