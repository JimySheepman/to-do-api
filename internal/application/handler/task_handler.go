package handler

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func CreateTask(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	task, err := domain.service.CreateTask(customContext)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   task,
	})
}

func ListTask(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	task, err := domain.service.ListTask(customContext)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   task,
	})
}

func UpdateTask(c *fiber.Ctx) error {

	return errors.New("")
}
func DeleteTask(c *fiber.Ctx) error {

	return errors.New("")
}
func CreateComment(c *fiber.Ctx) error {

	return errors.New("")
}
func ListComment(c *fiber.Ctx) error {

	return errors.New("")
}
func UpdateComment(c *fiber.Ctx) error {

	return errors.New("")
}
func DeleteComment(c *fiber.Ctx) error {

	return errors.New("")
}
