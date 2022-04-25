package service

import (
	"context"
	"time"

	"github.com/JimySheepman/to-do-api/internal/domain/model"
	"github.com/JimySheepman/to-do-api/internal/domain/repository"
)

type commentService struct {
	commentRepository repository.CommentRepository
}

func NewCommentService(r repository.CommentRepository) CommentService {
	return &commentService{
		commentRepository: r,
	}
}

func (s *commentService) CreateComment(ctx context.Context, comment *model.Comment) error {
	comment.CreatedAt = time.Now()
	return s.commentRepository.CreateComment(ctx, comment)
}

func (s *commentService) ListComment(ctx context.Context) (*[]model.Comment, error) {
	return s.commentRepository.ListComment(ctx)
}

func (s *commentService) UpdateComment(ctx context.Context, id int, comment *model.Comment) error {
	return s.commentRepository.UpdateComment(ctx, id, comment)
}

func (s *commentService) DeleteComment(ctx context.Context, id int) error {
	return s.commentRepository.DeleteComment(ctx, id)
}
