package handlers

import (
	"bandb/models"
	"bandb/src/forms"
	"bandb/src/helpers"
	"bandb/src/render"
	"log"
	"net/http"
	"strconv"
	"time"
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
		helpers.ClientError(w, http.StatusBadRequest)
		return
	}

	sd := r.Form.Get("start_date")
	ed := r.Form.Get("end_date")

	// go time quirk related
	timeLayout := "2006-01-02"
	startDate, err := time.Parse(timeLayout, sd)
	if err != nil {
		helpers.ClientError(w, http.StatusBadRequest)
	}
	endDate, err := time.Parse(timeLayout, ed)
	if err != nil {
		helpers.ClientError(w, http.StatusBadRequest)
	}

	roomId, err := strconv.Atoi(r.Form.Get("room_id"))
	if err != nil {
		helpers.ClientError(w, http.StatusBadRequest)
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
		StartDate: startDate,
		EndDate:   endDate,
		RoomID:    roomId,
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

	err = m.DB.InsertReservation(reservation)
	if err != nil {
		helpers.ClientError(w, http.StatusInternalServerError)
	}
	m.App.Session.Put(r.Context(), "reservation", reservation)

	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)

	if !ok {
		helpers.ClientError(w, http.StatusNotFound)
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
