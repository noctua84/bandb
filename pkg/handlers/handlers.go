package handlers

import (
	"bandb/models"
	"bandb/pkg/config"
	"bandb/pkg/render"
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

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	td := models.TemplateData{}

	td.Data = map[string]interface{}{
		"title":       "Home Page",
		"description": "Welcome to the home page",
	}

	render.UseTemplate(w, "home.page", &td)
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	td := models.TemplateData{}
	td.Data = map[string]interface{}{
		"title":       "About Us",
		"description": "This is the about page",
		"remote_ip":   remoteIP,
	}

	render.UseTemplate(w, "about.page", &td)
}
