package model

import "time"

type Commend struct {
	User        string    `json:"user"`
	TaskId      int       `json:"task_id"`
	UserComment string    `json:"comment"`
	CreatedAt   time.Time `json:"created_at"`
}
