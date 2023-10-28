package handlers

import "github.com/gofiber/fiber/v2"

type UserHandler struct{}

func (t *UserHandler) UserPongResponse(c *fiber.Ctx) error {
	return c.SendString("Pong from user")
}
