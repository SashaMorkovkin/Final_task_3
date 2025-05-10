package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			login TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL
		);
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			expression TEXT NOT NULL,
			result TEXT NOT NULL,
			user_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES users(id)
		);
	`)
	if err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}
}

func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Printf("Failed to close database: %v", err)
	}
}
