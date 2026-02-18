package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./jobs.db")

	if err != nil {
		log.Fatal(err)
	}

	schema := `
	CREATE TABLE IF NOT EXISTS jobs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		url TEXT NOT NULL,
		description TEXT,
		date_added DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err = db.Exec(schema)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
