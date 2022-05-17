package handler

import (
	"context"
	"strconv"

	"github.com/JimySheepman/to-do-api/internal/domain/comment"
	"github.com/JimySheepman/to-do-api/internal/service"
	"github.com/gofiber/fiber/v2"
)

type CommentHandler struct {
	commentService service.CommentService
}

func NewCommentHandler(commentRoute fiber.Router, service service.CommentService) {

	handler := &CommentHandler{
		commentService: service,
	}

	commentRoute.Post("/", handler.createComment)
	commentRoute.Get("/", handler.listComment)
	commentRoute.Put("/:id", handler.updateComment)
	commentRoute.Delete("/:id", handler.deleteComment)
}

func (h *CommentHandler) createComment(c *fiber.Ctx) error {
	comment := &comment.Comment{}

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

	return c.SendStatus(fiber.StatusCreated)
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
		"data": comments,
	})
}

func (h *CommentHandler) updateComment(c *fiber.Ctx) error {
	comment := &comment.Comment{}
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

	return c.SendStatus(fiber.StatusOK)
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
