package service

import (
	"context"
	"time"

	"github.com/JimySheepman/to-do-api/internal/domain/task"
)

type TaskService interface {
	CreateTask(ctx context.Context, task *task.Task) error
	ListTask(ctx context.Context) (*[]task.Task, error)
	UpdateTask(ctx context.Context, id int, task *task.Task) error
	DeleteTask(ctx context.Context, id int) error
}

type taskService struct {
	taskRepository task.TaskRepository
}

func NewTaskService(r task.TaskRepository) TaskService {
	return &taskService{
		taskRepository: r,
	}
}

func (s *taskService) CreateTask(ctx context.Context, task *task.Task) error {
	task.CreatedAt = time.Now()
	return s.taskRepository.CreateTask(ctx, task)
}

func (s *taskService) ListTask(ctx context.Context) (*[]task.Task, error) {
	return s.taskRepository.ListTask(ctx)
}

func (s *taskService) UpdateTask(ctx context.Context, id int, task *task.Task) error {
	return s.taskRepository.UpdateTask(ctx, id, task)
}

func (s *taskService) DeleteTask(ctx context.Context, id int) error {
	return s.taskRepository.DeleteTask(ctx, id)
}
