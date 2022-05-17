package task

import (
	"context"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task *Task) error
	ListTask(ctx context.Context) (*[]Task, error)
	UpdateTask(ctx context.Context, id int, task *Task) error
	DeleteTask(ctx context.Context, id int) error
}
