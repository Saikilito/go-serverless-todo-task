package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/saikilito/go-serverless-todo-task/app/task/src/routes"
)

var fiberLambda *fiberadapter.FiberLambda

// Handler
func health(c *fiber.Ctx) error {
	return c.SendString("Server UP 👋!")
}

var app *fiber.App
// init the Fiber Server
func init() {
	// log.Printf("Fiber cold start")
	app = fiber.New()

	// Midleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Routes
	app.Get("/", health)
	routes.MainRoutes(app)


	fiberLambda = fiberadapter.New(app)
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(fiberLambda.ProxyWithContextV2)
}