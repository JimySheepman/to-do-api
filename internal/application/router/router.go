package router

import (
	"github.com/JimySheepman/to-do-api/internal/application/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	task := api.Group("/task")
	task.Post("/create", handler.CreateTask)
	task.Get("/list", handler.ListTask)
	task.Patch("/update", handler.UpdateTask)
	task.Delete("/delete", handler.DeleteTask)

	comment := api.Group("/comment")
	comment.Post("/create", handler.CreateComment)
	comment.Get("/list", handler.ListComment)
	comment.Patch("/update", handler.UpdateComment)
	comment.Delete("/delete", handler.DeleteComment)
}
