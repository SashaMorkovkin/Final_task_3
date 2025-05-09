package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

// InitDB открывает соединение с базой данных и создаёт таблицы.
func InitDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := createTables(db); err != nil {
		db.Close()
		return nil, fmt.Errorf("error creating tables: %w", err)
	}

	log.Println("Database initialized successfully")
	return db, nil
}

// createTables создаёт необходимые таблицы, если они не существуют.
func createTables(db *sql.DB) error {
	userTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		login TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);`
	if _, err := db.Exec(userTable); err != nil {
		return err
	}

	tasksTable := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		expression TEXT NOT NULL,
		result TEXT NOT NULL,
		user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`
	if _, err := db.Exec(tasksTable); err != nil {
		return err
	}

	log.Println("Tables created or verified successfully")
	return nil
}
