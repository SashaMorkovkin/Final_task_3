package db

import (
	"database/sql"
	"log"
)

var DB *sql.DB

// Инициализация базы данных SQLite
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./calculator.db")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Создаем таблицы, если они не существуют
	createTables()
}

// Создание таблиц в базе данных
func createTables() {
	// Таблица пользователей
	query := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		login TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Error creating users table: %v", err)
	}

	// Таблица задач
	query = `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		expression TEXT NOT NULL,
		result TEXT NOT NULL,
		user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`
	_, err = DB.Exec(query)
	if err != nil {
		log.Fatalf("Error creating tasks table: %v", err)
	}
}

// Закрытие подключения к базе данных
func CloseDB() {
	err := DB.Close()
	if err != nil {
		log.Fatalf("Error closing database: %v", err)
	}
}
