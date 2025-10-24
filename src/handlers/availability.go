package handlers

import (
	"bandb/models"
	"bandb/src/render"
	"encoding/json"
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

// AvailabilityJSON handles availability and sends JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	type jsonResponse struct {
		OK      bool   `json:"ok"`
		Message string `json:"message"`
	}

	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(out)
	if err != nil {
		return
	}
}
