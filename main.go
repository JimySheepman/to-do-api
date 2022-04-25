package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JimySheepman/to-do-api/internal/application/handler"
	"github.com/JimySheepman/to-do-api/internal/domain/service"
	"github.com/JimySheepman/to-do-api/internal/infrastructure/persistence"
	"github.com/JimySheepman/to-do-api/internal/infrastructure/persistence/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

const idleTimeout = 10 * time.Second

func GracefulShutdown(app *fiber.App, port string) {
	// Listen from a different goroutine
	go func() {
		if err := app.Listen(port); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")
	fmt.Println("Fiber was successful shutdown.")
}

func main() {

	postgresql, err := persistence.ConnectDB()
	if err != nil {
		log.Fatal("Database connection error: $s", err)
	}

	app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
	})

	app.Use(cors.New())
	app.Use(logger.New())

	taskRepository := repository.NewTaskRepository(postgresql)
	commentRepository := repository.NewCommentRepository(postgresql)

	taskService := service.NewTaskService(taskRepository)
	commentService := service.NewCommentService(commentRepository)

	handler.NewTaskHandler(app.Group("/api/v1/task"), taskService)
	handler.NewCommentHandler(app.Group("/api/v1/comment"), commentService)

	app.All("*", func(c *fiber.Ctx) error {
		errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())

		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status":  "fail",
			"message": errorMessage,
		})
	})

	GracefulShutdown(app, ":8080")
}
