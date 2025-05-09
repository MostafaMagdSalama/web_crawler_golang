# Web Crawler Project Documentation

## Project Structure

```
web_crawler_golang/
├── generate_documentation.ipynb  # Jupyter Notebook for generating documentation
├── go.mod                        # Go module file
├── go.sum                        # Go dependencies checksum file
├── cmd/
│   └── main.go                   # Entry point of the application
├── internal/
│   ├── crawler/
│   │   └── crawler.go            # Core crawling logic
│   ├── fetcher/
│   │   └── fetcher.go            # Fetching and parsing HTML content
│   └── storage/
│       └── storage.go            # Storage logic for results
├── result/
│   └── results.json              # JSON file containing crawled URLs
```

## Purpose

The Web Crawler project is designed to recursively fetch and parse web pages starting from a given URL. It extracts links from the pages and stores the visited URLs in a structured format.

## Components

### 1. `cmd/main.go`
- **Purpose**: Acts as the entry point of the application.
- **Key Features**:
  - Initializes the crawler with a starting URL and depth.
  - Manages concurrency using `sync.WaitGroup`.
  - Writes the results to a JSON file.

### 2. `internal/crawler/crawler.go`
- **Purpose**: Implements the core crawling logic.
- **Key Features**:
  - Normalizes URLs to avoid duplicate visits.
  - Tracks visited URLs using a thread-safe map.
  - Recursively crawls links up to a specified depth.

### 3. `internal/fetcher/fetcher.go`
- **Purpose**: Handles fetching and parsing of HTML content.
- **Key Features**:
  - Uses an HTTP client with a timeout to fetch web pages.
  - Extracts valid links from the HTML content.

### 4. `internal/storage/storage.go`
- **Purpose**: Manages storage of crawled results.
- **Key Features**:
  - Writes the visited URLs to a JSON file.

### 5. `result/results.json`
- **Purpose**: Stores the URLs visited by the crawler.
- **Format**: JSON object where keys are URLs and values indicate whether they were successfully visited.

## How to Use

1. **Run the Application**:
   ```bash
   go run cmd/main.go
   ```
2. **Specify Starting URL and Depth**:
   - Modify the `startURL` and `depth` variables in `cmd/main.go`.
3. **View Results**:
   - Check the `result/results.json` file for the list of visited URLs.

## Future Enhancements

- Add support for filtering URLs based on domain or patterns.
- Implement a more robust error-handling mechanism.
- Optimize concurrency for large-scale crawling.

---

This documentation provides an overview of the Web Crawler project, its structure, and functionality. For further details, refer to the source code files.