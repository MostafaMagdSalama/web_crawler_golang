package fetcher

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/html"
)

type FetcherService interface {
	fetch(url string) []string
}

var client = &http.Client{
	Timeout: 10 * time.Second,
}

func isValidURL(link string) bool {
	parsedURL, err := url.Parse(link)
	if err != nil {
		return false
	}
	return parsedURL.Scheme == "http" || parsedURL.Scheme == "https"
}

func Fetch(url string) ([]string, error) {
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Failed:", url)
		return nil, err
	}

	defer resp.Body.Close()

	links := []string{}

	z := html.NewTokenizer(resp.Body)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}
		t := z.Token()
		if t.Type == html.StartTagToken && t.Data == "a" {
			for _, a := range t.Attr {
				if a.Key == "href" {
					if isValidURL(a.Val) {
						links = append(links, a.Val)
					}
				}
			}
		}
	}
	return links, nil
}
