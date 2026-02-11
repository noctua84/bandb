package render

import (
	"bandb/models"
	"html/template"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData
	r, err := requestWithSession()

	if err != nil {
		t.Fatal(err)
	}

	session.Put(r.Context(), "flash", "Test flash message")
	result := AddDefaultData(&td, r)

	if result.Flash != "Test flash message" {
		t.Error("AddDefaultData did not add flash message to TemplateData")
	}
}

func TestCreateTemplateCache(t *testing.T) {
	// Create temp directory with test templates
	tmpDir := t.TempDir()

	// Create a minimal layout
	layoutContent := `{{define "base"}}<!DOCTYPE html><html>{{template "content" .}}</html>{{end}}`
	err := os.WriteFile(filepath.Join(tmpDir, "base.layout.tmpl"), []byte(layoutContent), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Create partials directory
	partialsDir := filepath.Join(tmpDir, "partials")
	err = os.MkdirAll(partialsDir, 0755)
	if err != nil {
		t.Fatal(err)
	}

	// Create a minimal page
	pageContent := `{{template "base" .}}{{define "content"}}<h1>Test</h1>{{end}}`
	err = os.WriteFile(filepath.Join(tmpDir, "home.page.tmpl"), []byte(pageContent), 0644)
	if err != nil {
		t.Fatal(err)
	}

	cache := CreateTemplateCache(tmpDir)

	if cache == nil {
		t.Error("CreateTemplateCache returned nil")
	}

	if _, ok := cache["home.page"]; !ok {
		t.Error("Expected 'home.page' template in cache")
	}
}

func TestCreateTemplateCacheEmptyDir(t *testing.T) {
	tmpDir := t.TempDir()

	// Create partials subdirectory (required by Glob)
	err := os.MkdirAll(filepath.Join(tmpDir, "partials"), 0755)
	if err != nil {
		t.Fatal(err)
	}

	cache := CreateTemplateCache(tmpDir)

	if cache == nil {
		t.Error("CreateTemplateCache returned nil for empty directory")
	}

	if len(cache) != 0 {
		t.Errorf("Expected empty cache, got %d entries", len(cache))
	}
}

func TestUseTemplate(t *testing.T) {
	// Create temp directory with test templates
	tmpDir := t.TempDir()

	layoutContent := `{{define "base"}}<!DOCTYPE html><html>{{template "content" .}}</html>{{end}}`
	_ = os.WriteFile(filepath.Join(tmpDir, "base.layout.tmpl"), []byte(layoutContent), 0644)
	_ = os.MkdirAll(filepath.Join(tmpDir, "partials"), 0755)

	pageContent := `{{template "base" .}}{{define "content"}}<h1>Hello</h1>{{end}}`
	_ = os.WriteFile(filepath.Join(tmpDir, "test.page.tmpl"), []byte(pageContent), 0644)

	// Set up cache
	testApp.TemplateCache = CreateTemplateCache(tmpDir)
	testApp.UseCache = true

	rr := httptest.NewRecorder()
	req, err := requestWithSession()
	if err != nil {
		t.Fatal(err)
	}

	td := &models.TemplateData{}

	UseTemplate(rr, req, "test.page", td)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}

	if rr.Body.Len() == 0 {
		t.Error("Expected non-empty response body")
	}
}

func TestUseTemplateNotFound(t *testing.T) {
	testApp.TemplateCache = make(map[string]*template.Template)
	testApp.UseCache = true

	rr := httptest.NewRecorder()
	req, err := requestWithSession()
	if err != nil {
		t.Fatal(err)
	}

	UseTemplate(rr, req, "nonexistent", &models.TemplateData{})

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", rr.Code)
	}
}

func TestUseTemplateWithoutCacheControlled(t *testing.T) {
	// Save original working directory
	originalWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	// Create temp directory structure
	tmpDir := t.TempDir()
	templatesDir := filepath.Join(tmpDir, "templates")
	err = os.MkdirAll(filepath.Join(templatesDir, "partials"), 0755)
	if err != nil {
		t.Fatal(err)
	}

	layoutContent := `{{define "base"}}<!DOCTYPE html><html>{{template "content" .}}</html>{{end}}`
	_ = os.WriteFile(filepath.Join(templatesDir, "base.layout.tmpl"), []byte(layoutContent), 0644)

	pageContent := `{{template "base" .}}{{define "content"}}<h1>No Cache</h1>{{end}}`
	_ = os.WriteFile(filepath.Join(templatesDir, "nocache.page.tmpl"), []byte(pageContent), 0644)

	// Change to temp directory so "./templates" resolves correctly
	err = os.Chdir(tmpDir)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(originalWd)

	testApp.UseCache = false

	rr := httptest.NewRecorder()
	req, err := requestWithSession()
	if err != nil {
		t.Fatal(err)
	}

	UseTemplate(rr, req, "nocache.page", &models.TemplateData{})

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}

	if rr.Body.Len() == 0 {
		t.Error("Expected non-empty response body")
	}
}
