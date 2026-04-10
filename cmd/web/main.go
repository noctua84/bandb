package main

import (
	"bandb/models"
	"bandb/src/config"
	"bandb/src/driver"
	"bandb/src/handlers"
	"bandb/src/helpers"
	"bandb/src/render"
	"database/sql"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
)

// config
const port string = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}

	defer func(SQL *sql.DB) {
		err := SQL.Close()
		if err != nil {
			log.Fatal("Error closing database connection: " + err.Error())
		}
	}(db.SQL)

	fmt.Printf("Starting Server and listening on %s\n", port)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (*driver.DB, error) {
	gob.Register(models.Reservation{})
	gob.Register(models.Room{})
	gob.Register(models.User{})
	gob.Register(models.Restriction{})
	
	// change this to true in production
	app.InProduction = false

	app.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.ErrorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app.UseCache = false

	// session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// connect to database
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("postgresql://postgres:postgres@localhost:5432/bandb?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	log.Println("Connected to database")

	// create template cache
	if app.UseCache {
		fmt.Println("Creating template cache")

		tc := render.CreateTemplateCache("./templates")
		if tc == nil {
			return nil, fmt.Errorf("could not create template cache")
		}

		// assign template cache to app config
		app.TemplateCache = tc
		fmt.Println("Template cache: ", app.TemplateCache)
	}

	render.NewTemplates(&app)

	// handlers repository
	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	helpers.NewHelpers(&app)

	return db, nil
}
