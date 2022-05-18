package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	var err error

	dsn := fmt.Sprintf("postgres://postgres:root@localhost:5432/postgres?sslmode=disable")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db, nil
}
