package persistence

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_HOST     = "127.0.0.1"
	DB_PORT     = 5432
	DB_USER     = "postgres"
	DB_PASSWORD = "root"
	DB_NAME     = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	var err error
	/* 	TODO:every time .env file load error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	*/
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
