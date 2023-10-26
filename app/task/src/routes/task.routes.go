package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saikilito/go-serverless-todo-task/app/task/src/handlers"
)

func taskRoutes(app *fiber.App, handlers *handlers.TaskHandler) {
	taskRoutes := app.Group("/task")

	taskRoutes.Get("/ping", handlers.TaskPongResponse)
}