package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	endpoint "webCrawler/EndPoint"
)

var (
	checkServerInterval = 60 * time.Minute
	oldPagetime         = 60 // in minutes
)

func setupAPIEndpoints() {
	http.HandleFunc("/numWorkers", endpoint.NumWorkersHandler)
	http.HandleFunc("/speedPerHour", endpoint.SpeedPerHourHandler)
	http.HandleFunc("/crawl", endpoint.CrawlHandler)
}

// This function runs every hour to remove old pages in cache
func myHourlyFunction() {
	// fmt.Println("This function runs every hour.")
	for _, page := range endpoint.Pages {
		if page.Crosstime(float64(oldPagetime)) {
			fmt.Printf("delete a url %s\n", page.URL)
			delete(endpoint.Pages, page.URL)
		}
	}
}

func main() {
	setupAPIEndpoints()

	// Serve the HTML page.
	http.Handle("/", http.FileServer(http.Dir("./Static")))

	server := &http.Server{
		Addr:         ":8000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Create a quit channel to stop the function when the server is shut down.
	var quit = make(chan struct{})

	ticker := time.NewTicker(checkServerInterval)
	// Run the function in a goroutine.
	go func() {
		for {
			select {
			case <-ticker.C:
				myHourlyFunction()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	log.Printf("Server listening on :8000")
	log.Fatal(server.ListenAndServe())
	close(quit)
}
