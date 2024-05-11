package parsers

import (
	"net"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/html"
)

type linksParser struct {
}

func NewLinksParser() linksParser {
	return linksParser{}
}

func (l *linksParser) Parse(data string, host string, analysis AnalysisResponse) {

	tokenizer := html.NewTokenizer(strings.NewReader(data))
	for {
		switch tokenizer.Next() {
		case html.StartTagToken:
			token := tokenizer.Token()
			if token.Data == "a" {
				url := GetAttribute(token, "href")

				//Only add valid links
				if IsUrl(url) {
					linkObject := Link{
						Url:       url,
						LinkType:  getLinkType(url, host),
						Reachable: checkLinkAccessability(url),
					}
					analysis.AddLink(linkObject)
				}
			}

			if token.Data == "script" {
				url := GetAttribute(token, "src")

				//Only add valid links
				if IsUrl(url) {
					linkObject := Link{
						Url:       url,
						LinkType:  getLinkType(url, host),
						Reachable: checkLinkAccessability(url),
					}
					analysis.AddLink(linkObject)
				}
			}

			if token.Data == "link" {
				url := GetAttribute(token, "href")

				//Only add valid links
				if IsUrl(url) {
					linkObject := Link{
						Url:       url,
						LinkType:  getLinkType(url, host),
						Reachable: checkLinkAccessability(url),
					}
					analysis.AddLink(linkObject)
				}
			}
		case html.ErrorToken:
			return
		}
	}
}

// get content of an attribute given the tag
// returns empty array if attribute is not found
func GetAttribute(token html.Token, attribute string) string {
	for _, val := range token.Attr {
		if val.Key == attribute {
			return val.Val
		}
	}

	return ""
}

// get link type: internal or external
func getLinkType(link string, host string) string {
	if strings.Contains(link, host) {
		return "Internal"
	}
	return "External"
}

// validate url
func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

// check if url is accessible or not
func checkLinkAccessability(link string) bool {
	timeout := 1 * time.Second
	_, err := net.DialTimeout("tcp", link, timeout)
	return err != nil
}
