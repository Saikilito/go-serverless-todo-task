package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saikilito/go-serverless-todo-task/app/task/src/handlers"
)

func userRoutes(app *fiber.App, handlers *handlers.UserHandler) {
	userRoutes := app.Group("/user")

	userRoutes.Get("/ping", handlers.UserPongResponse)
}