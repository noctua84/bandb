package handlers

import (
	"bandb/models"
	"bandb/src/forms"
	"bandb/src/render"
	"log"
	"net/http"
)

// Reservation handles the reservation page (GET)
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation

	td := models.TemplateData{
		Data: map[string]interface{}{
			"title":       "Reservation",
			"description": "This is the reservation page",
			"reservation": emptyReservation,
		},
		Form: forms.New(nil),
	}

	render.UseTemplate(w, r, "reservation.page", &td)
}

// PostReservation handles the reservation page (POST)
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3, r)

	if !form.Valid() {
		td := models.TemplateData{
			Data: map[string]interface{}{
				"title":       "Reservation",
				"description": "This is the reservation page",
				"reservation": reservation,
			},
			Form: form,
		}

		render.UseTemplate(w, r, "reservation.page", &td)
		return
	}

	log.Printf("%+v\n", reservation)

	m.App.Session.Put(r.Context(), "reservation", reservation)

	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)

	if !ok {
		log.Println("reservation not found in session")
		m.App.Session.Put(r.Context(), "error", "reservation not found in session or reservation not valid")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")
	
	td := models.TemplateData{
		Data: map[string]interface{}{
			"reservation": reservation,
		},
	}

	render.UseTemplate(w, r, "reservation-summary.page", &td)
}
