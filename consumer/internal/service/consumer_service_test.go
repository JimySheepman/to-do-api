package service

import (
	"reflect"
	"testing"
	"time"

	"github.com/JimySheepman/to-do-api/consumer/internal/domain/comment"
	"github.com/segmentio/kafka-go"
)

type postgresqlMockCommentRepository struct {
	update map[string]error
}

func NewMockCommentRepository() postgresqlMockCommentRepository {
	return postgresqlMockCommentRepository{
		update: map[string]error{},
	}
}

func (r *postgresqlMockCommentRepository) OnUpdate(statu string, err error) {
	r.update[statu] = err
}

func (r *postgresqlMockCommentRepository) UpdateComment(createdAt time.Time, statu string, comment *comment.Comment) error {
	return r.update[statu]
}

func TestUpdateComment(t *testing.T) {
	r := NewMockCommentRepository()
	s := NewCommentService(&r)

	var tests = []struct {
		testName string
		statu    string
		createAt time.Time
		comment  *comment.Comment
		key      []byte
		value    []byte
		expect   error
	}{
		{
			"IsMessageKeyComment True",
			"",
			time.Now(),
			&comment.Comment{},
			[]byte("comment"),
			[]byte{123, 34, 110, 97, 109, 101, 34, 58, 34, 49, 34, 125, 10},
			nil,
		},
		{
			"IsMessageKeyComment False",
			"reject",
			time.Now(),
			&comment.Comment{},
			[]byte{34, 99, 111, 109, 109, 101, 110, 116, 34, 10},
			[]byte{123, 34, 110, 97, 109, 101, 34, 58, 34, 49, 34, 125, 10},
			nil,
		},
		{
			"IsMessageKeyComment False",
			"reject",
			time.Now(),
			&comment.Comment{},
			[]byte{34, 99, 111, 109, 109, 101, 110, 116, 34, 10},
			[]byte{},
			nil,
		},
	}

	for _, test := range tests {

		t.Run(test.testName, func(t *testing.T) {
			r.OnUpdate(test.statu, test.expect)
			actual := s.UpdateComment(kafka.Message{Key: test.key, Value: test.value})
			if !reflect.DeepEqual(actual, test.expect) {
				t.Errorf("%d value error", test.expect)
			}
		})
	}
}
