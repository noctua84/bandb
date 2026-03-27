package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// DB holds the database connection pool
type DB struct {
	SQL *sql.DB
}

var dbCon = &DB{}

// maybe move them to a config/.env in the future
const maxOpenConns = 10
const maxIdleConns = 5
const maxConnLifetime = 5 * time.Minute

// ConnectSQL create connection pool for Postgres
func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}

	d.SetMaxOpenConns(maxOpenConns)
	d.SetMaxIdleConns(maxIdleConns)
	d.SetConnMaxLifetime(maxConnLifetime)

	dbCon.SQL = d

	err = TestDB(d)
	if err != nil {
		return nil, err
	}

	return dbCon, nil
}

// NewDatabase creates a new database for the app
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// TestDB helper function to test the established connection.
func TestDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}

	return nil
}
