package service

import (
	"context"
	"time"

	"github.com/JimySheepman/to-do-api/internal/domain/model"
	"github.com/JimySheepman/to-do-api/internal/domain/repository"
)

type taskService struct {
	taskRepository repository.TaskRepository
}

func NewTaskService(r repository.TaskRepository) TaskService {
	return &taskService{
		taskRepository: r,
	}
}

func (s *taskService) CreateTask(ctx context.Context, task *model.Task) error {
	task.CreatedAt = time.Now()
	return s.taskRepository.CreateTask(ctx, task)
}

func (s *taskService) ListTask(ctx context.Context) (*[]model.Task, error) {
	return s.taskRepository.ListTask(ctx)
}

func (s *taskService) UpdateTask(ctx context.Context, id int, task *model.Task) error {
	return s.taskRepository.UpdateTask(ctx, id, task)
}

func (s *taskService) DeleteTask(ctx context.Context, id int) error {
	return s.taskRepository.DeleteTask(ctx, id)
}
