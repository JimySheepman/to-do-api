package repository

import (
	"database/sql"
	"log"
	"reflect"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/JimySheepman/to-do-api/consumer/internal/domain/comment"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestUpdateComment(t *testing.T) {
	db, mock := NewMock()
	repo := NewCommentRepository(db)

	comment := &comment.Comment{
		Id:          1,
		TaskId:      1,
		UserName:    "test",
		UserComment: "test",
		Statu:       "test",
		CreatedAt:   time.Now(),
	}

	t.Run("successful update", func(t *testing.T) {
		query := "UPDATE comments SET statu = \\$1 WHERE created_at = \\$2"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(comment.Statu, comment.CreatedAt).WillReturnResult(sqlmock.NewResult(0, 1))

		actual := repo.UpdateComment(comment.CreatedAt, comment.Statu, comment)
		var want error = nil
		if !reflect.DeepEqual(actual, want) {
			t.Errorf("got:%v expect:%v", actual, want)
		}
	})

	t.Run("ExecContext error", func(t *testing.T) {
		query := "UPDATE comments SET statu = \\$1 WHERE created_at = \\$2"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(comment.Statu, comment.CreatedAt).WillReturnResult(sqlmock.NewResult(0, 1))

		actual := repo.UpdateComment(comment.CreatedAt, comment.Statu, comment)
		var want error = nil
		if !reflect.DeepEqual(actual, want) {
			t.Errorf("got:%v expect:%v", actual, want)
		}
	})
}
