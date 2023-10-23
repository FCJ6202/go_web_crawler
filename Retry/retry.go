package retry

import (
	"html/template"
	"log"
	"net/http"
)

func RetryPage(w http.ResponseWriter, URL, payingStatus string) {
	tmpl, err := template.ParseFiles("Retry/retry.html")
	if err != nil {
		http.Error(w, "Failed to load retry page", http.StatusInternalServerError)
		log.Printf("Error : %s", err)
		return
	}

	data := struct {
		URL          string
		PayingStatus string
	}{
		URL:          URL,
		PayingStatus: payingStatus,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Failed to render retry page", http.StatusInternalServerError)
		log.Printf("Error : %s", err)
	}
}
