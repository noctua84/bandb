package handlers

import (
	"bandb/models"
	"bandb/pkg/render"
	"log"
	"net/http"
)

// Availability handles the availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	td := models.TemplateData{
		Data: map[string]interface{}{
			"title":       "Search for Availability",
			"description": "Search for available rooms by start and end dates",
		},
	}

	render.UseTemplate(w, r, "availability.page", &td)
}

// PostAvailability handles the post request for availability
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	startDate := r.FormValue("start_date")
	endDate := r.FormValue("end_date")

	log.Printf("Received availability request: Start Date: %s, End Date: %s", startDate, endDate)

	_, err := w.Write([]byte("Post request for availability received. Form data logged."))
	if err != nil {
		return
	}
}
