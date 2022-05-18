package comment

type CommentRepository interface {
	UpdateComment(id int, statu string, comment *Comment) error
}
