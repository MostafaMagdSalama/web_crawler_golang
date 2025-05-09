package main

import (
	"fmt"
	"sync"
	"web_crawler/internal/crawler"
	"web_crawler/internal/storage"
)

func main() {
	var wg sync.WaitGroup
	jobs := make(chan string)

	startURL := "https://www.axisapp.com/"
	depth := 3

	wg.Add(1)
	go func() {
		crawler.Crawl(startURL, &wg, jobs, depth)
		close(jobs)
	}()

	wg.Wait()

	if err := storage.WriteResultsToFile(crawler.VisitedURLs()); err != nil {
		fmt.Println("Error writing results to file:", err)
	}
}
