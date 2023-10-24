package endpoint

import (
	"fmt"
	"net/http"
	"strconv"
)

// It calls when admin give post or get request in /numWorkers endpoint.
// In get request admin will see the currently number of worker work in this system.
// In post request admin will be able to edit the number of worker in this system.
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
