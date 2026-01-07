package main

import (
	"bandb/src/config"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	if mux == nil {
		t.Error("routes() returned nil handler")
	}

	// Further type assertion to check if mux is of expected type
	switch m := mux.(type) {
	case *chi.Mux:
		// Test passed
	default:
		t.Errorf("routes() returned type %T, expected chi.Mux", m)
	}
}
