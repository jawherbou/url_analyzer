package main

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/net/html"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
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

		tokenizer := html.NewTokenizer(strings.NewReader(string(body)))

		var title string

	out:
		for {
			switch tokenizer.Next() {
			case html.StartTagToken:
				token := tokenizer.Token()
				if token.Data == "title" {
					tokenizer.Next()
					log.Println("Getting the title:", tokenizer.Token().Data)
					title = tokenizer.Token().Data
					break out
				}
			case html.ErrorToken:
				break out
			}
		}

		return c.JSON(fiber.Map{
			"title": title,
		})

	})

	app.Listen(":3000")
}
