package model

import "time"

type Task struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Category  string    `json:"category"`
	Statu     string    `json:"statu"`
	CreatedAt time.Time `json:"created_at"`
}
