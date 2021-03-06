package main

import (
	"log"

	"github.com/JimySheepman/to-do-api/consumer/internal/application"
	"github.com/JimySheepman/to-do-api/consumer/internal/infrastructure/consumer"
	"github.com/JimySheepman/to-do-api/consumer/internal/infrastructure/db"
	"github.com/JimySheepman/to-do-api/consumer/internal/infrastructure/db/repository"
	"github.com/JimySheepman/to-do-api/consumer/internal/service"
)

const (
	KAFKA_URL = "127.0.0.1"
	TOPIC     = "my-topic"
	GROUP_ID  = "0"
)

func main() {

	postgresql, err := db.ConnectDB("postgres", "postgres://postgres:root@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal("Database connection error: $s", err)
	}

	commentRepository := repository.NewCommentRepository(postgresql)

	commentService := service.NewCommentService(commentRepository)

	reader, err := consumer.GetKafkaReader(KAFKA_URL, TOPIC, GROUP_ID)
	if err != nil {
		log.Fatal(err)
	}

	consume := application.NewCommentConsume(commentService)
	consume.Consuming(reader)

}
