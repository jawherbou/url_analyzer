package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"github.com/jawherbou/url_analyzer/server/controllers"
)

func Test_GetAnalysis_ValidPath(t *testing.T) {
	app := fiber.New()

	// Define a test handler using the GetAnalysis function from the controller
	app.Get("/analyze", controllers.GetAnalysis)

	// Create a new HTTP request for testing
	req := httptest.NewRequest(http.MethodGet, "/analyze?url=https://example.com", nil)
	res, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	// Assert that the response status code is 200 OK
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_GetAnalysis_InvalidPath(t *testing.T) {
	app := fiber.New()

	// Define a test handler using the GetAnalysis function from the controller
	app.Get("/analyze", controllers.GetAnalysis)

	// Create a new HTTP request for testing
	req := httptest.NewRequest(http.MethodGet, "/analyze?url=example", nil)
	res, _ := app.Test(req, -1)

	// Assert that the response status code is 400
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}
