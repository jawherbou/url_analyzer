package controllers

import (
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/jawherbou/url_analyzer/server/services/parsers"
	"github.com/jawherbou/url_analyzer/server/services/urls"
)

func GetAnalysis(c *fiber.Ctx) error {

	// get url from query
	url := strings.TrimSpace(c.Query("url"))

	urlParser := urls.NewUrlInfo()
	urlInfo, err := urlParser.GetUrlInfo(url)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"err": err,
		})
	}

	body := urlInfo.Body
	host := urlInfo.Host

	// create report
	analysis := parsers.NewAnalysisResponse()

	// initiate parsers
	titleParser := parsers.NewTitleParser()
	htmlVersionParser := parsers.NewHtmlVersionParser()
	loginFormParser := parsers.NewLoginFormParser()
	headingsParser := parsers.NewHeadingsParser()
	linksParser := parsers.NewLinksParser()

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(5) // Number of goroutines

	// Run the function calls concurrently using goroutines
	go func() {
		defer wg.Done()
		titleParser.Parse(string(body), analysis)
	}()

	go func() {
		defer wg.Done()
		htmlVersionParser.Parse(string(body), analysis)
	}()

	go func() {
		defer wg.Done()
		loginFormParser.Parse(string(body), analysis)
	}()

	go func() {
		defer wg.Done()
		headingsParser.Parse(string(body), analysis)
	}()

	go func() {
		defer wg.Done()
		linksParser.Parse(string(body), host, analysis)
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	// Return status 200 OK.
	return c.JSON(analysis)
}
