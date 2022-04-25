package service

import (
	"context"

	"github.com/JimySheepman/to-do-api/internal/domain/model"
)

type CommentService interface {
	CreateComment(ctx context.Context, comment *model.Comment) error
	ListComment(ctx context.Context) (*[]model.Comment, error)
	UpdateComment(ctx context.Context, id int, comment *model.Comment) error
	DeleteComment(ctx context.Context, id int) error
}
