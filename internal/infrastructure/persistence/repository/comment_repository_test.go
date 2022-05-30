package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"reflect"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/JimySheepman/to-do-api/internal/domain/comment"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

// TODO: find how to fecth a PrepareContext err

func TestCreateComment(t *testing.T) {
	db, mock := NewMock()
	repo := NewCommentRepository(db)

	comment := &comment.Comment{
		TaskId:      1,
		UserName:    "test",
		UserComment: "test",
		Statu:       "test",
		CreatedAt:   time.Now(),
	}

	t.Run("succesful create", func(t *testing.T) {
		query := "INSERT INTO comments \\(task_id, user_name, user_comment, statu,created_at\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5\\)"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(comment.TaskId,
			comment.UserName,
			comment.UserComment,
			comment.Statu,
			comment.CreatedAt).WillReturnResult(sqlmock.NewResult(0, 1))

		actual := repo.CreateComment(context.Background(), comment)
		if !reflect.DeepEqual(actual, nil) {
			t.Errorf("got:%v expect:%v", actual, nil)
		}
	})

	t.Run("PrepareContext error", func(t *testing.T) {
		query := "INSERT INTO comments \\(task_id, user_name, user_comment, statu,created_at\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5\\)"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs().WillReturnResult(sqlmock.NewResult(0, 1))

		actual := repo.CreateComment(context.Background(), comment)
		want := errors.New("")
		if !reflect.DeepEqual(actual, want) {
			t.Errorf("got:%v expect:%v", actual, want)
		}
	})

	t.Run("ExecContext error", func(t *testing.T) {
		query := "INSERT INTO comments \\(task_id, user_name, user_comment, statu,created_at\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5\\)"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(comment.TaskId,
			comment.UserComment,
			comment.Statu,
			comment.CreatedAt).WillReturnResult(sqlmock.NewResult(0, 1))

		actual := repo.CreateComment(context.Background(), comment)
		want := errors.New("ExecQuery 'INSERT INTO comments (task_id, user_name, user_comment, statu,created_at) VALUES ($1, $2, $3, $4, $5)', arguments do not match: expected 4, but got 5 arguments")
		if !reflect.DeepEqual(actual, want) {
			t.Errorf("got:%v expect:%v", actual, want)
		}
	})
}

// TODO: find how to write a selec query mock test
func TestListComment(t *testing.T) {
	db, mock := NewMock()
	repo := NewCommentRepository(db)
	comment := &comment.Comment{
		Id:          1,
		TaskId:      1,
		UserName:    "test",
		UserComment: "test",
		Statu:       "approved",
		CreatedAt:   time.Now()}

	t.Run("query could not match actual sql", func(t *testing.T) {
		query := "SELECT * FROM comments"
		rows := sqlmock.NewRows([]string{"id", "task_id", "user_name", "user_comment", "statu", "created_at"})
		mock.ExpectQuery(query).WillReturnRows(rows)

		actual, _ := repo.ListComment(context.Background())
		var want error = nil
		if !reflect.DeepEqual(actual, want) {
			t.Errorf("got:%v expect:%v", actual, want)
		}
	})

	t.Run("query could not match actual sql 2", func(t *testing.T) {
		query := "SELECT * FROM comments WHERE statu= \\$1"
		rows := sqlmock.NewRows([]string{"id", "task_id", "ser_name", "user_comment", "statu", "created_at"}).
			AddRow(comment.Id, comment.TaskId, comment.UserName, comment.UserComment, comment.Statu, comment.CreatedAt)
		mock.ExpectQuery(query).WillReturnRows(rows)

		actual, _ := repo.ListComment(context.Background())
		var want error = nil
		if !reflect.DeepEqual(actual, want) {
			t.Errorf("got:%v expect:%v", actual, want)
		}
	})
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

	t.Run("succesful update", func(t *testing.T) {
		query := "UPDATE comments SET user_name = \\$1, user_comment = \\$2 WHERE id = \\$3"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(comment.UserName, comment.UserComment, comment.Id).WillReturnResult(sqlmock.NewResult(0, 1))

		actual := repo.UpdateComment(context.Background(), comment.Id, comment)
		var want error = nil
		if !reflect.DeepEqual(actual, want) {
			t.Errorf("got:%v expect:%v", actual, want)
		}
	})

	t.Run("ExecContext error", func(t *testing.T) {
		query := "UPDATE comments SET user_name = \\$1, user_comment = \\$2 WHERE id = \\$3"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(comment.UserName, comment.Id).WillReturnResult(sqlmock.NewResult(0, 1))

		actual := repo.UpdateComment(context.Background(), comment.Id, comment)
		var want error = nil
		if !reflect.DeepEqual(actual, want) {
			t.Errorf("got:%v expect:%v", actual, want)
		}
	})
}

func TestDeleteComment(t *testing.T) {
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

	t.Run("succesful delete", func(t *testing.T) {
		query := "DELETE FROM comments WHERE id = \\$1"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(comment.Id).WillReturnResult(sqlmock.NewResult(0, 1))

		actual := repo.DeleteComment(context.Background(), comment.Id)
		var want error = nil
		if !reflect.DeepEqual(actual, want) {
			t.Errorf("got:%v expect:%v", actual, want)
		}
	})

	t.Run("ExecContext error", func(t *testing.T) {
		query := "DELETE FROM comments WHERE id = \\$1"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(comment.UserName, comment.UserComment, comment.Id).WillReturnResult(sqlmock.NewResult(0, 1))

		actual := repo.DeleteComment(context.Background(), comment.Id)
		var want error = nil
		if !reflect.DeepEqual(actual, want) {
			t.Errorf("got:%v expect:%v", actual, want)
		}
	})

}
