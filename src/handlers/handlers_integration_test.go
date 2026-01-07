package handlers

import (
	"io"
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

	for _, tt := range tests {
		if tt.method == http.MethodGet {
			t.Run(tt.name, func(t *testing.T) {
				resp, err := client.Get(ts.URL + tt.path)
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
