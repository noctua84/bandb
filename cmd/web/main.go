package main

import (
	"bandb/pkg/config"
	"bandb/pkg/handlers"
	"bandb/pkg/render"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

// config
const port string = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	// change this to true in production
	app.InProduction = false

	app.UseCache = false

	// session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// create template cache
	if app.UseCache {
		fmt.Println("Creating template cache")

		tc := render.CreateTemplateCache()
		// assign template cache to app config
		app.TemplateCache = tc
		fmt.Println("Template cache: ", app.TemplateCache)
	}

	render.NewTemplates(&app)

	// handlers repository
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Printf("Starting Server and listening on %s\n", port)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
