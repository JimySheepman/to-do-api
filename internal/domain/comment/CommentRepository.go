package comment

import (
	"context"
)

type CommentRepository interface {
	CreateComment(ctx context.Context, comment *Comment) error
	ListComment(ctx context.Context) (*[]Comment, error)
	UpdateComment(ctx context.Context, id int, comment *Comment) error
	DeleteComment(ctx context.Context, id int) error
}
