package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAvailabilityHandlers(t *testing.T) {
	ts := httptest.NewServer(getTestRoutes())
	defer ts.Close()

	client := ts.Client()

	tests := []handlerTest{
		// GET tests
		{"Availability", "/availability", http.MethodGet, nil, http.StatusOK},

		// POST tests
		{"PostAvailability", "/availability", http.MethodPost, []postData{
			{key: "start_date", value: "2024-01-01"},
			{key: "end_date", value: "2024-01-07"},
		}, http.StatusOK},
		{"AvailabilityJSON", "/availability-json", http.MethodPost, nil, http.StatusOK},
	}

	runHandlerTests(t, client, ts.URL, tests)
}
