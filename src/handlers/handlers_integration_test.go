package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var tests = []handlerTest{
	{"Home", "/", http.MethodGet, nil, http.StatusOK},
	{"About", "/about", http.MethodGet, nil, http.StatusOK},
	{"Generals", "/generals", http.MethodGet, nil, http.StatusOK},
	{"Majors", "/majors", http.MethodGet, nil, http.StatusOK},
	{"Contact", "/contact", http.MethodGet, nil, http.StatusOK},
}

// Integration tests for handlers:
func TestHandlers(t *testing.T) {
	routes := getTestRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	// Client that accepts self-signed certs
	client := ts.Client()

	runHandlerTests(t, client, ts.URL, tests)
}
