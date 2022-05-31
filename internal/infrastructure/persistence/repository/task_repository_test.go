package repository

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/JimySheepman/to-do-api/internal/domain/task"
)

// TODO: find how to fecth a PrepareContext err
func TestCreateTask(t *testing.T) {
	db, mock := NewMock()
	repo := NewTaskRepository(db)

	task := &task.Task{
		Id:        1,
		Title:     "test",
		Content:   "test",
		Category:  "test",
		Statu:     "test",
		CreatedAt: time.Now(),
	}

	t.Run("successful create", func(t *testing.T) {
		query := "INSERT INTO tasks \\(title, content, category,statu,created_at\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5\\)"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(task.Title,
			task.Content,
			task.Category,
			task.Statu,
			task.CreatedAt).WillReturnResult(sqlmock.NewResult(0, 1))

		actual := repo.CreateTask(context.Background(), task)
		var want error = nil
		if !reflect.DeepEqual(actual, want) {
			t.Errorf("got:%v expect:%v", actual, want)
		}
	})

	t.Run("PrepareContext error", func(t *testing.T) {
		query := "INSERT INTO tasks \\(title, content, category,statu,created_at\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5\\)"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs().WillReturnResult(sqlmock.NewResult(0, 1))

		actual := repo.CreateTask(context.Background(), task)
		var want error = nil
		if !reflect.DeepEqual(actual, want) {
			t.Errorf("got:%v expect:%v", actual, want)
		}
	})

	t.Run("ExecContext error", func(t *testing.T) {
		query := "INSERT INTO tasks \\(title, content, category,statu,created_at\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5\\)"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(task.Title,
			task.Content,
			task.Category,
			task.CreatedAt).WillReturnResult(sqlmock.NewResult(0, 1))

		actual := repo.CreateTask(context.Background(), task)
		want := errors.New("ExecQuery 'INSERT INTO tasks (title, content, category,statu,created_at) VALUES ($1, $2, $3, $4, $5)', arguments do not match: expected 4, but got 5 arguments")
		if !reflect.DeepEqual(actual, want) {
			t.Errorf("got:%v expect:%v", actual, want)
		}
	})
}

// TODO: find how to write a selec query mock test
func TestListTask(t *testing.T) {
	db, mock := NewMock()
	repo := NewTaskRepository(db)

	task := &task.Task{
		Id:        1,
		Title:     "test",
		Content:   "test",
		Category:  "test",
		Statu:     "test",
		CreatedAt: time.Now(),
	}

	t.Run("query could not match actual sql", func(t *testing.T) {
		query := "SELECT * FROM tasks"
		rows := sqlmock.NewRows([]string{"id", "title", "content", "category", "statu", "created_at"})
		mock.ExpectQuery(query).WillReturnRows(rows)

		actual, _ := repo.ListTask(context.Background())
		var want error = nil
		if !reflect.DeepEqual(actual, want) {
			t.Errorf("got:%v expect:%v", actual, want)
		}
	})

	t.Run("query could not match actual sql 2", func(t *testing.T) {
		query := "SELECT * FROM tasks"
		rows := sqlmock.NewRows([]string{"id", "title", "content", "category", "statu", "created_at"}).
			AddRow(task.Id, task.Title, task.Content, task.Category, task.Statu, task.CreatedAt)
		mock.ExpectQuery(query).WillReturnRows(rows)

		actual, _ := repo.ListTask(context.Background())
		var want error = nil
		if !reflect.DeepEqual(actual, want) {
			t.Errorf("got:%v expect:%v", actual, want)
		}
	})
}

func TestUpdateTask(t *testing.T) {
	db, mock := NewMock()
	repo := NewTaskRepository(db)

	task := &task.Task{
		Id:        1,
		Title:     "test",
		Content:   "test",
		Category:  "test",
		Statu:     "test",
		CreatedAt: time.Now(),
	}

	t.Run("successful update", func(t *testing.T) {
		query := "UPDATE tasks SET title = \\$1, content = \\$2, category = \\$3, statu = \\$4 WHERE id = \\$5"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(task.Title, task.Content, task.Category, task.Statu, task.Id).WillReturnResult(sqlmock.NewResult(0, 1))

		actual := repo.UpdateTask(context.Background(), task.Id, task)
		var want error = nil
		if !reflect.DeepEqual(actual, want) {
			t.Errorf("got:%v expect:%v", actual, want)
		}
	})

	t.Run("ExecContext error", func(t *testing.T) {
		query := "UPDATE tasks SET title = \\$1, content = \\$2, category = \\$3, statu = \\$4 WHERE id = \\$5"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(task.Title, task.Content).WillReturnResult(sqlmock.NewResult(0, 1))

		actual := repo.UpdateTask(context.Background(), task.Id, task)
		var want = errors.New("ExecQuery 'UPDATE tasks SET title = $1, content = $2, category = $3, statu = $4 WHERE id = $5', arguments do not match: expected 2, but got 5 arguments")
		if !reflect.DeepEqual(actual, want) {
			t.Errorf("got:%v expect:%v", actual, want)
		}
	})
}

func TestDeleteTask(t *testing.T) {
	db, mock := NewMock()
	repo := NewTaskRepository(db)

	task := &task.Task{
		Id:        1,
		Title:     "test",
		Content:   "test",
		Category:  "test",
		Statu:     "test",
		CreatedAt: time.Now(),
	}

	t.Run("successful delete", func(t *testing.T) {
		query := "DELETE FROM tasks WHERE id = \\$1"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(task.Id).WillReturnResult(sqlmock.NewResult(0, 1))

		actual := repo.DeleteTask(context.Background(), task.Id)
		var want error = nil
		if !reflect.DeepEqual(actual, want) {
			t.Errorf("got:%v expect:%v", actual, want)
		}
	})

	t.Run("ExecContext error", func(t *testing.T) {
		query := "DELETE FROM tasks WHERE id = \\$1"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(task.Title, task.Content, task.Category).WillReturnResult(sqlmock.NewResult(0, 1))

		actual := repo.DeleteTask(context.Background(), task.Id)
		var want = errors.New("ExecQuery 'DELETE FROM tasks WHERE id = $1', arguments do not match: expected 3, but got 1 arguments")
		if !reflect.DeepEqual(actual, want) {
			t.Errorf("got:%v expect:%v", actual, want)
		}
	})

}
