package handlers

import (
	"bandb/models"
	"bandb/src/config"
	"bandb/src/render"
	"encoding/gob"
	"io"
	"net/http"
	"testing"

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

// -----------------------------------------
// Test setup
// -----------------------------------------
// init registers types for session storage
func init() {
	// Register types that will be stored in session
	gob.Register(models.Reservation{})
}

// getTestRepository creates a test repository with a session and template cache
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

// getTestRoutes sets up the routes for testing
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

// -----------------------------------------
// Test helpers
// -----------------------------------------
// runGetTest performs a GET request and checks the response status code
func runGetTest(t *testing.T, client *http.Client, baseURL string, tt handlerTest) {
	t.Helper()
	resp, err := client.Get(baseURL + tt.path)
	if err != nil {
		t.Fatalf("failed to make request: %v", err)
	}

	defer closeBody(t, resp.Body)

	if resp.StatusCode != tt.expectedStatus {
		t.Errorf("path %s: got %d, want %d", tt.path, resp.StatusCode, tt.expectedStatus)
	}
}

// runPostTest performs a POST request with form data and checks the response status code
func runPostTest(t *testing.T, client *http.Client, baseURL string, tt handlerTest) {
	t.Helper()
	formData := make(map[string][]string)
	for _, pd := range tt.params {
		formData[pd.key] = []string{pd.value}
	}

	resp, err := client.PostForm(baseURL+tt.path, formData)
	if err != nil {
		t.Fatalf("failed to make request: %v", err)
	}

	defer closeBody(t, resp.Body)

	if resp.StatusCode != tt.expectedStatus {
		t.Errorf("path %s: got %d, want %d", tt.path, resp.StatusCode, tt.expectedStatus)
	}
}

// runHandlerTests executes all tests in the slice using the appropriate method
func runHandlerTests(t *testing.T, client *http.Client, baseURL string, tests []handlerTest) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.method {
			case http.MethodGet:
				runGetTest(t, client, baseURL, tt)
			case http.MethodPost:
				runPostTest(t, client, baseURL, tt)
			}
		})
	}
}

// closeBody closes the response body and reports any error
func closeBody(t *testing.T, body io.ReadCloser) {
	t.Helper()
	if err := body.Close(); err != nil {
		t.Errorf("failed to close response body: %v", err)
	}
}
