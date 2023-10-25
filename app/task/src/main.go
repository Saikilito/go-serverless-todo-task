package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"

	"github.com/saikilito/go-serverless-todo-task/app/task/src/routes"
)

var fiberLambda *fiberadapter.FiberLambda

// Handler
func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

// init the Fiber Server
func init() {
	log.Printf("Fiber cold start")
	var app *fiber.App
	app = fiber.New()

	// Routes
	app.Get("/", hello)

	routes.MainRoutes(app)

	fiberLambda = fiberadapter.New(app)
}

// Handler will deal with Fiber working with Lambda
func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return fiberLambda.ProxyWithContextV2(ctx, req)
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(Handler)
}