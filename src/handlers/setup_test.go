package handlers

import (
	"bandb/models"
	"bandb/src/config"
	"bandb/src/render"
	"encoding/gob"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
)

type postData struct {
	key   string
	value string
}

type handlerTest = struct {
	name           string
	path           string
	method         string
	params         []postData
	expectedStatus int
}

func init() {
	// Register types that will be stored in session
	gob.Register(models.Reservation{})
}

func getTestRepository() *Repository {
	session := scs.New()
	app := &config.AppConfig{
		Session:      session,
		UseCache:     true,
		InProduction: false,
	}

	// Initialize the render package with the app config
	render.NewTemplates(app)

	// Create and cache templates for tests
	tc := render.CreateTemplateCache("../../templates")
	app.TemplateCache = tc

	return NewRepo(app)
}

func getTestRoutes() http.Handler {
	repo := getTestRepository()

	mux := chi.NewRouter()

	// Apply session middleware - this is required for handlers that use sessions
	mux.Use(repo.App.Session.LoadAndSave)

	mux.Get("/", repo.Home)
	mux.Get("/about", repo.About)
	mux.Get("/generals", repo.Generals)
	mux.Get("/majors", repo.Majors)
	mux.Get("/contact", repo.Contact)
	mux.Get("/reservation", repo.Reservation)
	mux.Post("/reservation", repo.PostReservation)
	mux.Get("/reservation-summary", repo.ReservationSummary)

	mux.Get("/availability", repo.Availability)
	mux.Post("/availability", repo.PostAvailability)
	mux.Post("/availability-json", repo.AvailabilityJSON)

	return mux
}
