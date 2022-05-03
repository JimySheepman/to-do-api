package handler

import (
	"context"
	"strconv"

	"github.com/JimySheepman/to-do-api/internal/domain/model"
	"github.com/JimySheepman/to-do-api/internal/domain/service"
	"github.com/gofiber/fiber/v2"
)

type TaskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(taskRoute fiber.Router, service service.TaskService) {
	handler := &TaskHandler{
		taskService: service,
	}

	taskRoute.Post("/create", handler.createTask)
	taskRoute.Get("/list", handler.listTask)
	taskRoute.Put("/update/:id", handler.updateTask)
	taskRoute.Delete("/delete/:id", handler.deleteTask)
}

func (h *TaskHandler) createTask(c *fiber.Ctx) error {
	task := &model.Task{}

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

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "Task has been created successfully!",
	})
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
		"status": "success",
		"data":   tasks,
	})
}

func (h *TaskHandler) updateTask(c *fiber.Ctx) error {
	task := &model.Task{}
	paramsMap := c.AllParams()

	targetedTaskId, err := strconv.Atoi(paramsMap["id"])
	if err != nil {
		return err
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
		return err
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
