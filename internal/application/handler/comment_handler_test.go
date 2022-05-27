package handler

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/JimySheepman/to-do-api/internal/domain/comment"
	"github.com/JimySheepman/to-do-api/internal/infrastructure/broker/producer"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type commentMockService struct {
	update map[int]error
}

func NewCommentMockService() *commentMockService {
	return &commentMockService{
		update: make(map[int]error),
	}
}

func (s *commentMockService) OnUpdate(id int, err error) {
	s.update[id] = err
}

func (s *commentMockService) CreateComment(ctx context.Context, comment *comment.Comment) error {
	comment.CreatedAt = time.Now()
	producer.Send("comment", comment)
	return nil
}

func (s *commentMockService) ListComment(ctx context.Context) (*[]comment.Comment, error) {
	return nil, nil
}

func (s *commentMockService) UpdateComment(ctx context.Context, id int, comment *comment.Comment) error {
	return s.update[id]
}

func (s *commentMockService) DeleteComment(ctx context.Context, id int) error {
	return nil
}

func Test_updateComment_Success(t *testing.T) {
	s := NewCommentMockService()
	//setup
	h := newCommentHandler(s)
	a := fiber.New()
	ctx := a.AcquireCtx(&fasthttp.RequestCtx{})

	//gıven
	c := &comment.Comment{}

	_, err := json.Marshal(c)

	if err != nil {
		t.Fail()
	}

	//ctx.Params("id", "1100")
	// ctx.Request().AppendBody(arr)

	//when
	s.OnUpdate(1100, nil)

	actual := h.updateComment(ctx)
	var expect error

	if actual != expect {
		t.Error()
	}

}
