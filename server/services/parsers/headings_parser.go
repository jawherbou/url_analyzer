package parsers

import (
	"log"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

type headingsParser struct {
}

func NewHeadingsParser() headingsParser {
	return headingsParser{}
}

func (h *headingsParser) Parse(data string, analysis AnalysisResponse) {

	regex, err := regexp.Compile("[hH][1-9]")
	if err != nil {
		log.Println("Error in compiling the regex", err)
		return
	}

	tokenizer := html.NewTokenizer(strings.NewReader(data))
	for {
		switch tokenizer.Next() {
		case html.StartTagToken:
			token := tokenizer.Token()
			match := regex.Match([]byte(token.Data))
			if match {
				tokenizer.Next()
				analysis.AddHeading(token.Data, tokenizer.Token().Data)
			}
		case html.ErrorToken:
			return
		}
	}
}
