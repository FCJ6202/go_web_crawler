package endpoint

import (
	"fmt"
	"net/http"
	"strconv"
)

func NumWorkersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		numWorkersStr := r.FormValue("workers")
		newWorker, err := strconv.Atoi(numWorkersStr)
		if err != nil {
			http.Error(w, "Invalid number of workers", http.StatusBadRequest)
			return
		}

		numWorkers = newWorker
	} else if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Current number of workers: %d", numWorkers)
	}
}
