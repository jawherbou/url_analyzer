package parsers

import (
	"strings"

	"golang.org/x/net/html"
)

type titleParser struct {
}

func NewTitleParser() titleParser {
	return titleParser{}
}

func (tp *titleParser) Parse(data string, analysis AnalysisResponse) {

	tokenizer := html.NewTokenizer(strings.NewReader(data))
	for {
		switch tokenizer.Next() {
		case html.StartTagToken:
			token := tokenizer.Token()
			if token.Data == "title" {
				tokenizer.Next()
				analysis.SetTitle(tokenizer.Token().Data)
				return
			}
		case html.ErrorToken:
			return
		}
	}
}
