package repository

import (
	"context"
	"database/sql"

	"github.com/JimySheepman/to-do-api/internal/domain/model"
	"github.com/JimySheepman/to-do-api/internal/domain/repository"
)

const (
	QUERY_CREATE_TASK = "INSERT INTO tasks (title, content, category,statu,created_at) VALUES (?, ?, ?, ?, ?)"
	QUERY_GET_TASKS   = "SELECT * FROM tasks"
	QUERY_UPDATE_TASK = "UPDATE tasks SET title = ?, content = ?, category = ?, statu = ? WHERE id = ?"
	QUERY_DELETE_TASK = "DELETE FROM tasks WHERE id = ?"
)

type postgresqlTaskRepository struct {
	potgresql *sql.DB
}

func NewTaskRepository(postgresqlConnection *sql.DB) repository.TaskRepository {
	return &postgresqlTaskRepository{
		potgresql: postgresqlConnection,
	}
}

func (r *postgresqlTaskRepository) CreateTask(ctx context.Context, task *model.Task) error {

	stmt, err := r.potgresql.PrepareContext(ctx, QUERY_CREATE_TASK)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, task.Title, task.Content, task.Category, task.Statu, task.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresqlTaskRepository) ListTask(ctx context.Context) (*[]model.Task, error) {
	var tasks []model.Task
	res, err := r.potgresql.QueryContext(ctx, QUERY_GET_TASKS)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	for res.Next() {
		task := &model.Task{}
		err = res.Scan(&task.Id, &task.Title, &task.Content, &task.Category, &task.Statu, &task.CreatedAt)
		if err != nil && err == sql.ErrNoRows {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, *task)
	}

	return &tasks, nil
}

func (r *postgresqlTaskRepository) UpdateTask(ctx context.Context, id int, task *model.Task) error {
	stmt, err := r.potgresql.PrepareContext(ctx, QUERY_UPDATE_TASK)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, task.Title, task.Content, task.Category, task.Statu, task.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresqlTaskRepository) DeleteTask(ctx context.Context, id int) error {
	stmt, err := r.potgresql.PrepareContext(ctx, QUERY_DELETE_TASK)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
