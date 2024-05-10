package controllers

import (
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jawherbou/url_analyzer/server/services/parsers"
)

func GetAnalysis(c *fiber.Ctx) error {

	// get url from query
	url := c.Query("url")

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error occurred when getting the response", err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"err": err,
		})
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error occurred when reading the response", err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"err": err,
		})
	}

	// create report
	analysis := parsers.NewAnalysisResponse()

	// get text
	titleParser := parsers.NewTitleParser()
	titleParser.Parse(string(body), analysis)

	// get html version
	htmlVersionParser := parsers.NewHtmlVersionParser()
	htmlVersionParser.Parse(string(body), analysis)

	// set login form status
	loginFormParser := parsers.NewLoginFormParser()
	loginFormParser.Parse(string(body), analysis)

	log.Println(analysis)

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
	})
}
