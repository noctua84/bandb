package render

import (
	"bandb/models"
	"bandb/pkg/config"
	"bytes"
	"github.com/justinas/nosurf"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, req *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(req)

	return td
}

func CreateTemplateCache() map[string]*template.Template {
	var cache = map[string]*template.Template{}

	// load layouts:
	layouts, err := filepath.Glob("templates/*.layout.tmpl")
	if err != nil {
		log.Fatalf("Error loading layout templates: %v", err)
	}

	partials, err := filepath.Glob("templates/partials/*.tmpl")
	if err != nil {
		log.Fatalf("Error loading partial templates: %v", err)
	}

	pages, err := filepath.Glob("templates/*.page.tmpl")
	if err != nil {
		log.Fatalf("Error loading page templates: %v", err)
	}

	if len(layouts) == 0 && len(pages) == 0 {
		log.Println("Warning: No templates found in templates directory")
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ext := filepath.Ext(name)
		t := name[:len(name)-len(ext)]

		allFiles := append([]string{page}, layouts...)
		allFiles = append(allFiles, partials...)

		tmpl, err := template.ParseFiles(allFiles...)

		if err != nil {
			log.Fatalf("Error parsing template %s: %v", t, err)
		}

		cache[t] = tmpl
	}

	return cache
}

func UseTemplate(w http.ResponseWriter, req *http.Request, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc = CreateTemplateCache()
	}

	t, ok := tc[tmpl]

	log.Printf("Using template %s", tmpl)

	if !ok {
		log.Printf("Template %s not found in cache", tmpl)
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, req)

	err := t.Execute(buf, td)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template %s: %v", tmpl, err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error writing template %s to response: %v", tmpl, err)
		return
	}
}
