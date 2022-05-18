package repository

import (
	"database/sql"
	"time"

	"github.com/JimySheepman/to-do-api/consumer/internal/domain/comment"
)

const (
	QUERY_UPDATE_COMMENT = "UPDATE comments SET statu = $1 WHERE created_at = $2"
)

type postgresqlCommentRepository struct {
	potgresql *sql.DB
}

func NewCommentRepository(postgresqlConnection *sql.DB) comment.CommentRepository {
	return &postgresqlCommentRepository{
		potgresql: postgresqlConnection,
	}
}

func (r *postgresqlCommentRepository) UpdateComment(createdAt time.Time, statu string, comment *comment.Comment) error {
	stmt, err := r.potgresql.Prepare(QUERY_UPDATE_COMMENT)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(statu, createdAt)
	if err != nil {
		return err
	}

	return nil
}
