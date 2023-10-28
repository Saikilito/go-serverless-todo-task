package handlers

import "github.com/gofiber/fiber/v2"

type TaskHandler struct{}

func (t *TaskHandler) TaskPongResponse(c *fiber.Ctx) error {
	return c.SendString("Pong from task")
}
