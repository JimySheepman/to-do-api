package repository

import (
	"database/sql"
)

// Queries that we will use.
const (
	QUERY_GET_USERS   = "SELECT * FROM users"
	QUERY_GET_USER    = "SELECT * FROM users WHERE id = ?"
	QUERY_CREATE_USER = "INSERT INTO users (name, address, created, modified) VALUES (?, ?, ?, ?)"
	QUERY_UPDATE_USER = "UPDATE users SET name = ?, address = ?, modified = ? WHERE id = ?"
	QUERY_DELETE_USER = "DELETE FROM users WHERE id = ?"
)

type potgresqlRepository struct {
	potgresql *sql.DB
}

func NewUserRepository(mariaDBConnection *sql.DB) Repository {
	return &potgresql{
		mariadb: mariaDBConnection,
	}
}
