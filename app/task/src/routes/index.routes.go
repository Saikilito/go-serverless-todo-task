package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saikilito/go-serverless-todo-task/app/task/src/handlers"
)

func MainRoutes(app *fiber.App) {
	taskRoutes(app, new(handlers.TaskHandler))
	userRoutes(app, new(handlers.UserHandler))
}