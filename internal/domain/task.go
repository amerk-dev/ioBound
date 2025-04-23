package domain

import "time"

type Task struct {
	ID        int        `json:"id"`
	Status    TaskStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"` //nil
}

type TaskStatus string

const (
	StatusReady TaskStatus = "ready"
	StatusWait  TaskStatus = "wait"
)
