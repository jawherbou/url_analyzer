package main

import (
	"log"

	"github.com/jawherbou/url_analyzer/server/middlewares"
	"github.com/jawherbou/url_analyzer/server/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Middlewares
	middlewares.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.

	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("Error starting Fiber server: %v", err)
	}
}
