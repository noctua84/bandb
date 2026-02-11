package handlers

import (
	"bandb/src/config"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Unit tests for handlers:
func TestNewRepo(t *testing.T) {
	app := &config.AppConfig{}
	repo := NewRepo(app)

	if repo.App != app {
		t.Error("expected App to be set in Repository")
	}
}

func TestNewHandlers(t *testing.T) {
	repo := &Repository{}
	NewHandlers(repo)

	if Repo != repo {
		t.Error("expected Repo to be set globally")
	}
}

func TestHome(t *testing.T) {
	repo := getTestRepository()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx, _ := repo.App.Session.Load(req.Context(), "")
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()

	repo.Home(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Home handler returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}
}

func TestAbout(t *testing.T) {
	repo := getTestRepository()

	req := httptest.NewRequest(http.MethodGet, "/about", nil)
	ctx, _ := repo.App.Session.Load(req.Context(), "")
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()

	repo.About(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("About handler returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}
}

func TestGenerals(t *testing.T) {
	repo := getTestRepository()

	req := httptest.NewRequest(http.MethodGet, "/generals", nil)
	ctx, _ := repo.App.Session.Load(req.Context(), "")
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()

	repo.Generals(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Generals handler returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}
}

func TestMajors(t *testing.T) {
	repo := getTestRepository()

	req := httptest.NewRequest(http.MethodGet, "/majors", nil)
	ctx, _ := repo.App.Session.Load(req.Context(), "")
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()

	repo.Majors(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Majors handler returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}
}

func TestContact(t *testing.T) {
	repo := getTestRepository()

	req := httptest.NewRequest(http.MethodGet, "/contact", nil)
	ctx, _ := repo.App.Session.Load(req.Context(), "")
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()

	repo.Contact(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Contact handler returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}
}
