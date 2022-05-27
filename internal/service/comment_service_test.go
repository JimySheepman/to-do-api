package service

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/JimySheepman/to-do-api/internal/domain/comment"
)

type postgresqlMockCommentRepository struct {
	create map[int]error
	list   []comment.Comment
	update map[int]error
	delete map[int]error
}

func NewMockCommentRepository() *postgresqlMockCommentRepository {
	return &postgresqlMockCommentRepository{
		create: map[int]error{},
		list:   []comment.Comment{},
		update: map[int]error{},
		delete: map[int]error{},
	}
}

func (r *postgresqlMockCommentRepository) OnCreate(id int, err error) {
	r.create[id] = err
}

func (r *postgresqlMockCommentRepository) OnList(c int) {
	r.list = make([]comment.Comment, c, c)
}

func (r *postgresqlMockCommentRepository) OnUpdate(id int, err error) {
	r.update[id] = err
}

func (r *postgresqlMockCommentRepository) OnDelete(id int, err error) {
	r.delete[id] = err
}

func (r *postgresqlMockCommentRepository) CreateComment(ctx context.Context, comment *comment.Comment) error {
	return r.create[comment.Id]
}

func (r *postgresqlMockCommentRepository) ListComment(ctx context.Context) (*[]comment.Comment, error) {
	return &r.list, nil
}

func (r *postgresqlMockCommentRepository) UpdateComment(ctx context.Context, id int, comment *comment.Comment) error {
	return r.update[id]
}

func (r *postgresqlMockCommentRepository) DeleteComment(ctx context.Context, id int) error {
	return r.delete[id]
}

func TestCreateComment(t *testing.T) {
	r := NewMockCommentRepository()
	s := NewCommentService(r)

	var tests = []struct {
		num    int
		tsk    *comment.Comment
		expect error
	}{
		{1, &comment.Comment{Id: 1}, nil},
		{3, &comment.Comment{Id: 3}, nil},
		{5, &comment.Comment{Id: 5}, nil},
		{10, &comment.Comment{}, errors.New("10")},
		{100, &comment.Comment{Id: 100}, errors.New("100")},
		{1000, &comment.Comment{Id: 1000}, errors.New("1000")},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("actual:%d expected error:%v", test.num, test.expect)
		t.Run(testName, func(t *testing.T) {
			r.OnCreate(test.num, test.expect)
			actual := s.CreateComment(context.Background(), test.tsk)
			if !reflect.DeepEqual(actual, test.expect) {
				t.Errorf("%d value error", test.num)
			}
		})
	}
}

func TestListComment(t *testing.T) {
	r := NewMockCommentRepository()
	s := NewCommentService(r)

	var tests = []struct {
		num    int
		expect error
	}{
		{1, nil},
		{2, nil},
		{3, nil},
		{4, nil},
		{5, nil},
		{6, nil},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("actual:%d expected error:%v", test.num, test.expect)
		t.Run(testName, func(t *testing.T) {
			r.OnList(test.num)
			actual, _ := s.ListComment(context.Background())
			if !reflect.DeepEqual(len(*actual), test.num) {
				t.Errorf("%d value error", test.num)
			}
		})
	}
}

func TestUpdateComment(t *testing.T) {
	r := NewMockCommentRepository()
	s := NewCommentService(r)

	var tests = []struct {
		num    int
		tsk    *comment.Comment
		expect error
	}{
		{1, &comment.Comment{}, nil},
		{3, &comment.Comment{}, nil},
		{5, &comment.Comment{}, nil},
		{10, &comment.Comment{}, errors.New("10")},
		{100, &comment.Comment{}, errors.New("100")},
		{1000, &comment.Comment{}, errors.New("1000")},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("actual:%d expected error:%v", test.num, test.expect)
		t.Run(testName, func(t *testing.T) {
			r.OnUpdate(test.num, test.expect)
			actual := s.UpdateComment(context.Background(), test.num, test.tsk)
			if !reflect.DeepEqual(actual, test.expect) {
				t.Errorf("%d value error", test.num)
			}
		})
	}
}

func TestDeleteComment(t *testing.T) {
	r := NewMockCommentRepository()
	s := NewCommentService(r)

	var tests = []struct {
		num    int
		tsk    *comment.Comment
		expect error
	}{
		{1, &comment.Comment{}, nil},
		{3, &comment.Comment{}, nil},
		{5, &comment.Comment{}, nil},
		{10, &comment.Comment{}, errors.New("10")},
		{100, &comment.Comment{}, errors.New("100")},
		{1000, &comment.Comment{}, errors.New("1000")},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("actual:%d expected error:%v", test.num, test.expect)
		t.Run(testName, func(t *testing.T) {
			r.OnDelete(test.num, test.expect)
			actual := s.DeleteComment(context.Background(), test.num)
			if !reflect.DeepEqual(actual, test.expect) {
				t.Errorf("%d value error", test.num)
			}
		})
	}
}
