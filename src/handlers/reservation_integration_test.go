package handlers

import (
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"
)

var testsReservation = []handlerTest{
	{"Reservation", "/reservation", http.MethodGet, nil, http.StatusOK},
	{"PostReservation", "/reservation", http.MethodPost, []postData{
		{key: "first_name", value: "John"},
		{key: "last_name", value: "Doe"},
		{key: "email", value: "john.doe@example.com"},
		{key: "phone", value: "123-456-7890"},
		{key: "start_date", value: "2024-10-01"},
		{key: "end_date", value: "2024-10-05"},
		{key: "room_id", value: "1"},
	}, http.StatusSeeOther},
	// Test for missing session data leading to redirect:
	{"ReservationSummary", "/reservation-summary", http.MethodGet, nil, http.StatusTemporaryRedirect},
}

// Integration tests for reservation handlers:
// These tests check the reservation-related handlers including session handling.
// ReservationSummary is only tested for redirect here, as it requires session data set by PostReservation.
func TestReservationHandlers(t *testing.T) {
	routes := getTestRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()
	// Client that accepts self-signed certs
	client := ts.Client()
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	for _, tt := range testsReservation {
		if tt.method == http.MethodGet {
			t.Run(tt.name, func(t *testing.T) {
				runGetTest(t, client, ts.URL, tt)
			})
		} else if tt.method == http.MethodPost {
			t.Run(tt.name, func(t *testing.T) {
				formData := make(map[string][]string)
				for _, pd := range tt.params {
					formData[pd.key] = []string{pd.value}
				}

				resp, err := client.PostForm(ts.URL+tt.path, formData)
				if err != nil {
					t.Fatalf("failed to make request: %v", err)
				}
				defer func(Body io.ReadCloser) {
					err := Body.Close()
					if err != nil {
						t.Errorf("failed to close response body: %v", err)
					}
				}(resp.Body)

				if resp.StatusCode != tt.expectedStatus {
					t.Errorf("path %s: got %d, want %d", tt.path, resp.StatusCode, tt.expectedStatus)
				}
			})
		}
	}
}

// Test the full reservation flow with session persistence
// This test simulates a user making a reservation and then viewing the summary
func TestReservationFlow(t *testing.T) {
	routes := getTestRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	jar, _ := cookiejar.New(nil)
	client := ts.Client()
	client.Jar = jar
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	// POST reservation
	formData := map[string][]string{
		"first_name": {"John"},
		"last_name":  {"Doe"},
		"email":      {"john.doe@example.com"},
		"phone":      {"123-456-7890"},
	}
	resp, err := client.PostForm(ts.URL+"/reservation", formData)
	if err != nil {
		t.Fatalf("POST failed: %v", err)
	}
	resp.Body.Close()

	// GET summary with session
	resp, err = client.Get(ts.URL + "/reservation-summary")
	if err != nil {
		t.Fatalf("GET failed: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Errorf("failed to close response body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("got %d, want %d", resp.StatusCode, http.StatusOK)
	}
}
