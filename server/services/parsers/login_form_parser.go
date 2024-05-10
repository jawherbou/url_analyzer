package parsers

import (
	"strings"

	"golang.org/x/net/html"
)

type loginFormParser struct {
}

func NewLoginFormParser() loginFormParser {
	return loginFormParser{}
}

func (lf *loginFormParser) Parse(data string, analysis AnalysisResponse) {

	tokenizer := html.NewTokenizer(strings.NewReader(data))
	for {
		switch tokenizer.Next() {
		case html.StartTagToken:
			token := tokenizer.Token()
			if token.Data == "form" {
				// if we have form tag, check for next tags if we have input password
				for {
					switch tokenizer.Next() {
					case html.StartTagToken, html.SelfClosingTagToken:
						token := tokenizer.Token()
						// if we have input tag, check if it is a password type input
						if token.Data == "input" {
							for _, val := range token.Attr {
								if val.Key == "type" && val.Val == "password" {
									analysis.SetHasLogin(true)
									return
								}
							}
						}
					case html.ErrorToken:
						return
					}
				}
			}
		case html.ErrorToken:
			return
		}
	}
}
