package comment

import "time"

type CommentRepository interface {
	UpdateComment(createdAt time.Time, statu string, comment *Comment) error
}
