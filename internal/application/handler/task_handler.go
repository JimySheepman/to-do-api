package handler

import (
	"context"
	"strconv"

	"github.com/JimySheepman/to-do-api/internal/domain/task"
	"github.com/JimySheepman/to-do-api/internal/service"
	"github.com/gofiber/fiber/v2"
)

type TaskHandler struct {
	taskService service.TaskService
}

func NewTaskRouter(taskRoute fiber.Router, service service.TaskService) {
	handler := newTaskHandler(service)

	taskRoute.Post("/", handler.createTask)
	taskRoute.Get("/", handler.listTask)
	taskRoute.Put("/:id", handler.updateTask)
	taskRoute.Delete("/:id", handler.deleteTask)
}

func newTaskHandler(service service.TaskService) *TaskHandler {
	handler := &TaskHandler{
		taskService: service,
	}

	return handler
}

func (h *TaskHandler) createTask(c *fiber.Ctx) error {
	task := &task.Task{}

	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	err := h.taskService.CreateTask(customContext, task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *TaskHandler) listTask(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	tasks, err := h.taskService.ListTask(customContext)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"data": tasks,
	})
}

func (h *TaskHandler) updateTask(c *fiber.Ctx) error {
	task := &task.Task{}
	paramsMap := c.AllParams()

	targetedTaskId, err := strconv.Atoi(paramsMap["id"])
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	err = h.taskService.UpdateTask(customContext, targetedTaskId, task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Task has been updated successfully!",
	})
}

func (h *TaskHandler) deleteTask(c *fiber.Ctx) error {
	paramsMap := c.AllParams()

	targetedTaskId, err := strconv.Atoi(paramsMap["id"])
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = h.taskService.DeleteTask(customContext, targetedTaskId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
