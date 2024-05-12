package controllers

import (
	"strings"

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

	// get text
	titleParser.Parse(string(body), analysis)

	// get html version
	htmlVersionParser.Parse(string(body), analysis)

	// set login form status
	loginFormParser.Parse(string(body), analysis)

	// add headings
	headingsParser.Parse(string(body), analysis)

	// add links
	linksParser.Parse(string(body), host, analysis)

	// Return status 200 OK.
	return c.JSON(analysis)
}
