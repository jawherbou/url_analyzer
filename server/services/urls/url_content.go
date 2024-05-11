package urls

import (
	"io"
	"log"
	"net/http"
)

type UrlInfo struct {
	Body []byte
	Host string
}

func NewUrlInfo() UrlInfo {
	return UrlInfo{}
}

func (u *UrlInfo) GetUrlInfo(url string) (*UrlInfo, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error occurred when getting the response", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error occurred when reading the response", err)
		return nil, err
	}

	urlInfo := &UrlInfo{
		Body: body,
		Host: resp.Request.Host,
	}

	return urlInfo, nil
}
