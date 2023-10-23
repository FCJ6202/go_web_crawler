package endpoint

import (
	"fmt"
	"net/http"
	"strconv"
)

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
