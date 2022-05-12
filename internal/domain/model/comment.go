package model

import "time"

type Comment struct {
	Id          int       `json:"id"`
	TaskId      int       `json:"task_id"`
	UserName    string    `json:"user_name"`
	UserComment string    `json:"user_comment"`
	CreatedAt   time.Time `json:"created_at"`
}