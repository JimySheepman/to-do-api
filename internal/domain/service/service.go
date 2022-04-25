package service

import (
	"context"
)

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) CreateTaskService(ctx context.Context) (*[]User, error) {
	return s.Repository.GetUsers(ctx)
}
