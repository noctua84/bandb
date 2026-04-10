package dbrepo

import (
	"bandb/src/config"
	"bandb/src/repository"
	"database/sql"
)

type postgresRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(app *config.AppConfig, con *sql.DB) repository.DatabaseRepo {
	return &postgresRepo{
		App: app,
		DB:  con,
	}
}
