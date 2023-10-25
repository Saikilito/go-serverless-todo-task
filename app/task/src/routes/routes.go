package routes

import "github.com/gofiber/fiber/v2"

func helloRoute(e *fiber.App){
	e.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello from route")
	})
}

func MainRoutes(app *fiber.App) {
	helloRoute(app)
}