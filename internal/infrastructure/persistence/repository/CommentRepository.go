package repository

import (
	"context"
	"database/sql"

	"github.com/JimySheepman/to-do-api/internal/domain/model"
	"github.com/JimySheepman/to-do-api/internal/domain/repository"
)

const (
	QUERY_CREATE_COMMENT = "INSERT INTO comments (task_id, user_name, user_comment,created_at) VALUES ($1, $2, $3, $4)"
	QUERY_GET_COMMENTS   = "SELECT * FROM comments"
	QUERY_UPDATE_COMMENT = "UPDATE comments SET username = $1, user_comment = $2 WHERE id = $3"
	QUERY_DELETE_COMMENT = "DELETE FROM comments WHERE id = $1"
)

type postgresqlCommentRepository struct {
	potgresql *sql.DB
}

func NewCommentRepository(postgresqlConnection *sql.DB) repository.CommentRepository {
	return &postgresqlCommentRepository{
		potgresql: postgresqlConnection,
	}
}

func (r *postgresqlCommentRepository) CreateComment(ctx context.Context, comment *model.Comment) error {

	stmt, err := r.potgresql.PrepareContext(ctx, QUERY_CREATE_COMMENT)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, comment.TaskId, comment.UserName, comment.UserComment, comment.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresqlCommentRepository) ListComment(ctx context.Context) (*[]model.Comment, error) {
	var commnets []model.Comment
	res, err := r.potgresql.QueryContext(ctx, QUERY_GET_COMMENTS)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	for res.Next() {
		commnet := &model.Comment{}
		err = res.Scan(&commnet.Id, &commnet.TaskId, &commnet.UserName, &commnet.UserComment, &commnet.CreatedAt)
		if err != nil && err == sql.ErrNoRows {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}
		commnets = append(commnets, *commnet)
	}

	return &commnets, nil
}

func (r *postgresqlCommentRepository) UpdateComment(ctx context.Context, id int, comment *model.Comment) error {
	stmt, err := r.potgresql.PrepareContext(ctx, QUERY_UPDATE_COMMENT)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, comment.UserName, comment.UserComment, comment.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresqlCommentRepository) DeleteComment(ctx context.Context, id int) error {
	stmt, err := r.potgresql.PrepareContext(ctx, QUERY_DELETE_COMMENT)
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
