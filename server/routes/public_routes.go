package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jawherbou/url_analyzer/server/controllers"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET method: we have only one
	route.Get("/", controllers.GetAnalysis) // get analysis for a webpage
}
