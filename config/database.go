package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./university.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable()
}

func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS grades (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		student_id TEXT,
		homework REAL,
		midterm REAL,
		final REAL,
		total REAL,
		grade TEXT
	);`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
