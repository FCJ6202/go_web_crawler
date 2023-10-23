package endpoint

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	model "webCrawler/Model"
	retry "webCrawler/Retry"
)

var (
	numPayWorker    = 5                                    // Worker for payed user
	numNonPayWorker = 2                                    // Worker for non payed user
	numWorkers      = 10                                   // Total number of worker
	speedPerHour    = 100                                  // number of page crawl per Hour
	oldPagetime     = 60                                   // in minutes
	payedWorker     = make(chan struct{}, numPayWorker)    // channel for payed worker
	nonpayedWorker  = make(chan struct{}, numNonPayWorker) // channel for unpayed worker
)

var Pages = make(map[string]*model.Page) // It contain previous data

// It convert the link in valid http format
func CovertValidLink(url string) string {
	if strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://") {
		return url
	} else {
		return fmt.Sprintf("http://%s", url)
	}
}

// It crawl the data from url using http.get API call and return response
func crawler(url string, w http.ResponseWriter) {
	// Check if the URL has been crawled in the last 60 minutes.
	value, ok := Pages[url]
	if ok && !value.Crosstime(float64(oldPagetime)) {
		fmt.Fprintf(w, "<p>%s</p>", value.Content.Data)
		return
	}

	if ok {
		delete(Pages, url)
	}

	// If the page is not cached, crawl it in real-time.
	res, err := http.Get(url)
	if err != nil {
		http.Error(w, "Crawl function Failed to fetch the page", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	// Read the data from the io.ReadCloser.
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	// You can store the content on the disk or cache it for future use here.
	Pages[url] = model.NewPage(url, string(data))

	// fmt.Println(content)
	// Return the content to the user.
	fmt.Fprintf(w, "<p>%s</p>", Pages[url].Content.Data)

	/* This code busy this function for 60s
	This line of code not for development this is only used in testing phase
	*/
	timeout := time.After(1 * 60 * time.Second)
	<-timeout

	fmt.Println(url)
}

// It call when user give post request in /crawl location
func CrawlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		// this is for Great
		if len(Pages) >= speedPerHour {
			http.Error(w, "Hourly crawl limit exceeded", http.StatusInternalServerError)
			return
		}

		// this is for required
		URL := r.FormValue("url")
		payingStatus := r.FormValue("pay") // 1 -> pay ,0 -> nonpay

		if payingStatus == "1" {
			select {
			case payedWorker <- struct{}{}:
				crawler(CovertValidLink(URL), w)
				<-payedWorker
			default:
				retry.RetryPage(w, URL)
			}
		} else {
			select {
			case nonpayedWorker <- struct{}{}:
				crawler(CovertValidLink(URL), w)
				<-nonpayedWorker
			default:
				retry.RetryPage(w, URL)
			}
		}
	} else {
		http.Error(w, "Only want post request", http.StatusInternalServerError)
		return
	}
}
