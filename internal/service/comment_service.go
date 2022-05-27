package service

import (
	"context"
	"time"

	"github.com/JimySheepman/to-do-api/internal/domain/comment"
	"github.com/JimySheepman/to-do-api/internal/infrastructure/broker/producer"
)

type CommentService interface {
	CreateComment(ctx context.Context, comment *comment.Comment) error
	ListComment(ctx context.Context) (*[]comment.Comment, error)
	UpdateComment(ctx context.Context, id int, comment *comment.Comment) error
	DeleteComment(ctx context.Context, id int) error
}

type commentService struct {
	commentRepository comment.CommentRepository
}

func NewCommentService(r comment.CommentRepository) CommentService {
	return &commentService{
		commentRepository: r,
	}
}

func (s *commentService) CreateComment(ctx context.Context, comment *comment.Comment) error {
	comment.CreatedAt = time.Now()
	// TODO: how to check producer error in service layer
	producer.Send("comment", comment)
	return s.commentRepository.CreateComment(ctx, comment)
}

func (s *commentService) ListComment(ctx context.Context) (*[]comment.Comment, error) {
	return s.commentRepository.ListComment(ctx)
}

func (s *commentService) UpdateComment(ctx context.Context, id int, comment *comment.Comment) error {
	return s.commentRepository.UpdateComment(ctx, id, comment)
}

func (s *commentService) DeleteComment(ctx context.Context, id int) error {
	return s.commentRepository.DeleteComment(ctx, id)
}
