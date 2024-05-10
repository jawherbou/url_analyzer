package main

import (
	"github.com/jawherbou/url_analyzer/server/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Middlewares

	// Routes
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.

	app.Listen(":3000")
}