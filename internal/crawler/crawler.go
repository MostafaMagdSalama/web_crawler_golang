package crawler

import (
	"fmt"
	"net/url"
	"sync"
	"web_crawler/internal/fetcher"
)

var visited = struct {
	m map[string]bool
	sync.Mutex
}{m: make(map[string]bool)}

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	parsedURL.Fragment = ""
	return parsedURL.String(), nil
}

func Crawl(url string, wg *sync.WaitGroup, jobs chan string, depth int) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	normalizedURL, err := normalizeURL(url)
	if err != nil {
		fmt.Println("Error normalizing URL:", url, "Error:", err)
		return
	}

	visited.Lock()
	if visited.m[normalizedURL] {
		visited.Unlock()
		return
	}
	visited.m[normalizedURL] = true
	visited.Unlock()

	fmt.Println("Visiting:", normalizedURL)
	links, err := fetcher.Fetch(normalizedURL)
	if err != nil {
		fmt.Println("Error fetching URL:", normalizedURL, "Error:", err)
		return
	}

	for _, link := range links {
		wg.Add(1)
		go Crawl(link, wg, jobs, depth-1)
	}
}

func VisitedURLs() map[string]bool {
	visited.Lock()
	defer visited.Unlock()

	copy := make(map[string]bool)
	for url, visited := range visited.m {
		copy[url] = visited
	}
	return copy
}
