package handlers

import (
	"bandb/models"
	"bandb/src/config"
	"bandb/src/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// Handler functions for the applications pages

// Home handles the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	td := models.TemplateData{}

	td.Data = map[string]interface{}{
		"title":       "Fort Tranquility B&B",
		"description": "Welcome Fort Tranquility B&B",
	}

	render.UseTemplate(w, r, "home.page", &td)
}

// About handles the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	td := models.TemplateData{}
	td.Data = map[string]interface{}{
		"title":       "About Us",
		"description": "This is the about page",
	}

	render.UseTemplate(w, r, "about.page", &td)
}

// Generals handles the generals quarters page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	td := models.TemplateData{
		Data: map[string]interface{}{
			"title":       "Generals Quarters",
			"description": "This is the generals quarters page",
		},
	}

	render.UseTemplate(w, r, "generals.page", &td)
}

// Majors handles the majors page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	td := models.TemplateData{
		Data: map[string]interface{}{
			"title":       "Majors Suite",
			"description": "This is the majors suite page",
		},
	}
	render.UseTemplate(w, r, "majors.page", &td)
}

// Contact handles the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	td := models.TemplateData{
		Data: map[string]interface{}{
			"title":       "Contact Us",
			"description": "This is the contact page",
		},
	}

	render.UseTemplate(w, r, "contact.page", &td)
}
