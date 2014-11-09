package scraper

import (
	"fmt"
	"net/http"
)

var (
	userAgents = map[string]string{
		"firefox": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:32.0) Gecko/20100101 Firefox/32.0",
		"iphone":  "Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25",
		"ipad":    "Mozilla/5.0 (iPad; CPU OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25",
		"android": "Mozilla/5.0(Linux; U; Android 2.2; en-gb; LG-P500 Build/FRF91) AppleWebKit/533.0 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
	}
)

type Getter func(url, agent string) *http.Response

type Scraper struct {
	getResponse Getter
}

func setupClient() http.Client {
	return http.Client{}
}

func NewScraper(g Getter) *Scraper {
	return &Scraper{getResponse: g}
}

func getSimpleResponse(url string) *http.Response {
	// user := setupClient()
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error with request %v", err)
	}
	return res
}

func redirectFunc(*http.Request, []*http.Request) error {
	return nil
}

func getResponse(url string, userAgent string) *http.Response {
	client := &http.Client{
		CheckRedirect: redirectFunc,
	}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", userAgents[userAgent])
	resp, _ := client.Do(req)
	return resp
}
