# go_web_crawler
A simple web crawler built using Go, much like a simplified Google, takes a link and provides the corresponding page associated with that link.

## Getting Started

- Clone the repo or download the zip file, and `cd` into the project folder.

```
cd go_web_crawler
go run main.go
```

- After the server starts, you have to open any browser and type a URL http://localhost:8000/

### Go Packages

```bash
/webCrawler/            # Contains web crawler logic.
/webCrawler/model       # Contains all struct definitions (Page, Content)
/webCrawler/endpoint    # Contain all endpoint handle function (/crawl,/numWorkers,/speedPerHour)
/webCrawler/retry       # Contain retry page logic
```

### High level Design
- drive link :- https://drive.google.com/file/d/1F7416viLlUJLeoldWbWZz5c_hWd8ysrZ/view?usp=sharing
