package service

import (
	"encoding/json"
	"strings"

	blacklist "github.com/JimySheepman/to-do-api/consumer/internal/domain/black_list"
	"github.com/JimySheepman/to-do-api/consumer/internal/domain/comment"
	kafka "github.com/segmentio/kafka-go"
)

const COMMENT = "comment"

type CommentService interface {
	UpdateComment(msg kafka.Message) error
}

type commentService struct {
	commentRepository comment.CommentRepository
}

func NewCommentService(r comment.CommentRepository) CommentService {
	return &commentService{
		commentRepository: r,
	}
}

func (s *commentService) UpdateComment(msg kafka.Message) error {
	if IsMessageKeyComment(msg.Key) {
		newMessgae, err := Unmarshal(msg.Value)
		if err != nil {
			return err
		}

		statu := ValidateMessageContent(newMessgae)

		return s.commentRepository.UpdateComment(newMessgae.Id, statu, newMessgae)
	}
	return nil
}

func Unmarshal(msg []byte) (*comment.Comment, error) {
	comment := &comment.Comment{}

	err := json.Unmarshal(msg, comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func IsMessageKeyComment(msg []byte) bool {
	if string(msg) == COMMENT {
		return true
	}
	return false
}

func ValidateMessageContent(msg *comment.Comment) string {
	for _, word := range blacklist.BlackList {
		res := strings.Contains(msg.UserComment, word)
		if res {
			return "reject"
		}
	}
	return "approved"
}
