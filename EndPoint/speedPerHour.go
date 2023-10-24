package endpoint

import (
	"fmt"
	"net/http"
	"strconv"
)

// It calls when admin give post or get request in /speedPerHour endpoint.
// In get request admin will see the currently how many page workers read in one hour in this system.
// In post request admin will be able to edit number of page read in one hour in this system.
func SpeedPerHourHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		speedPerHourStr := r.FormValue("speed")
		newSpeedPerHour, err := strconv.Atoi(speedPerHourStr)
		if err != nil {
			http.Error(w, "Invalid crawling speed", http.StatusBadRequest)
			return
		}

		speedPerHour = newSpeedPerHour
	} else if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Current crawling speed per hour: %d", speedPerHour)
	}
}
