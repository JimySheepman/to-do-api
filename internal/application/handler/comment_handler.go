package handler

import (
	"context"
	"strconv"

	"github.com/JimySheepman/to-do-api/internal/domain/model"
	"github.com/JimySheepman/to-do-api/internal/domain/service"
	"github.com/gofiber/fiber/v2"
)

type CommentHandler struct {
	commentService service.CommentService
}

func NewCommentHandler(commentRoute fiber.Router, service service.CommentService) {

	handler := &CommentHandler{
		commentService: service,
	}

	commentRoute.Post("/create", handler.createComment)
	commentRoute.Get("/list", handler.listComment)
	commentRoute.Put("/update/:id", handler.updateComment)
	commentRoute.Delete("/delete/:id", handler.deleteComment)
}

func (h *CommentHandler) createComment(c *fiber.Ctx) error {
	comment := &model.Comment{}

	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := c.BodyParser(comment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	err := h.commentService.CreateComment(customContext, comment)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "Comment has been created successfully!",
	})
}

func (h *CommentHandler) listComment(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	comments, err := h.commentService.ListComment(customContext)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   comments,
	})
}

func (h *CommentHandler) updateComment(c *fiber.Ctx) error {
	comment := &model.Comment{}
	paramsMap := c.AllParams()

	targetedId, err := strconv.Atoi(paramsMap["id"])
	if err != nil {
		return err
	}

	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := c.BodyParser(comment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	err = h.commentService.UpdateComment(customContext, targetedId, comment)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Comment has been updated successfully!",
	})
}

func (h *CommentHandler) deleteComment(c *fiber.Ctx) error {
	paramsMap := c.AllParams()

	targetedId, err := strconv.Atoi(paramsMap["id"])
	if err != nil {
		return err
	}

	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = h.commentService.DeleteComment(customContext, targetedId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
