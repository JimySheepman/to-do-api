package handler

import (
	"log"
	"testing"

	"github.com/JimySheepman/to-do-api/internal/domain/model"
	"github.com/JimySheepman/to-do-api/internal/domain/service"
	"github.com/JimySheepman/to-do-api/internal/infrastructure/persistence"
	"github.com/JimySheepman/to-do-api/internal/infrastructure/persistence/repository"
	"github.com/gofiber/fiber/v2"
)

func TestNewTaskHandler(t *testing.T) {
	postgresql, err := persistence.ConnectDB()
	if err != nil {
		log.Fatal("Database connection error: $s", err)
	}

	app := fiber.New()
	taskRepository := repository.NewTaskRepository(postgresql)
	taskService := service.NewTaskService(taskRepository)

	NewTaskHandler(app.Group("/api/v1/task"), taskService)
}

func TestListTask(t *testing.T) {
	var c *fiber.Ctx
	var tasks *[]model.Task
	taskHandler := TaskHandler{}

	got := taskHandler.ListTask(c)

	want := c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   tasks,
	})

	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
