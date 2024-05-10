package parsers

import (
	"strings"

	"golang.org/x/net/html"
)

type htmlVersionParser struct {
	types map[string]string
}

func NewHtmlVersionParser() htmlVersionParser {
	obj := htmlVersionParser{types: make(map[string]string)}
	obj.types["HTML 2.0"] = `<!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML 2.0//EN">`
	obj.types["HTML 3.2"] = `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 3.2 Final//EN">`
	obj.types["HTML 4.01 Strict"] = `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN">`
	obj.types["HTML 4.01 Transitional"] = `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">`
	obj.types["HTML 4.01 Frameset"] = `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Frameset//EN">`
	obj.types["XHTML 1.0 Strict"] = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">`
	obj.types["XHTML 1.0 Transitional"] = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/- xhtml1-transitional.dtd">`
	obj.types["XHTML 1.0 Frameset"] = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Frameset//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-frameset.dtd">`
	obj.types["XHTML 1.1"] = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN" "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">`
	obj.types["HTML 5"] = `<!DOCTYPE html>`

	return obj
}

func (h *htmlVersionParser) Parse(data string, analysis AnalysisResponse) {
	tokenizer := html.NewTokenizer(strings.NewReader(data))
	for {
		switch tokenizer.Next() {
		case html.DoctypeToken:
			token := tokenizer.Token()
			version := "Unknown"
			for name, value := range h.types {
				if token.String() == value {
					version = name
					break
				}
			}
			analysis.SetHtmlVersion(version)
			return
		case html.ErrorToken:
			return
		}
	}
}
