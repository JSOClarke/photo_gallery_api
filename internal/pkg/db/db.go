package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Underscore means that its a sideeffect so its not actually called but is used from an import that is called the db one above
)

// Connect to the database

func Connect(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
