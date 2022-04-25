package repository

import (
	"context"

	"github.com/JimySheepman/to-do-api/internal/domain/model"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task *model.Task) error
	ListTask(ctx context.Context) (*[]model.Task, error)
	UpdateTask(ctx context.Context, id int, task *model.Task) error
	DeleteTask(ctx context.Context, id int) error
}
