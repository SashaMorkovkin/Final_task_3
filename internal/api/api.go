package api

import (
	"database/sql"
	"fmt"
	"github.com/SashaMorkovkin/Final_task_3/internal/db"
)

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Task struct {
	ID         int    `json:"id"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
	UserID     int    `json:"user_id"`
}

// Функция для регистрации пользователя
func RegisterUser(login, password string) (*User, error) {
	// Проверка на существование пользователя
	var existingUser User
	err := db.DB.QueryRow("SELECT id, login FROM users WHERE login = ?", login).Scan(&existingUser.ID, &existingUser.Login)
	if err == nil {
		return nil, fmt.Errorf("user already exists")
	}

	// Добавляем нового пользователя в базу данных
	stmt, err := db.DB.Prepare("INSERT INTO users (login, password) VALUES (?, ?)")
	if err != nil {
		return nil, fmt.Errorf("could not prepare statement: %v", err)
	}
	_, err = stmt.Exec(login, password)
	if err != nil {
		return nil, fmt.Errorf("could not insert user: %v", err)
	}

	// Возвращаем нового пользователя
	return &User{Login: login, Password: password}, nil
}

// Функция для авторизации пользователя
func AuthenticateUser(login, password string) (*User, error) {
	var user User
	err := db.DB.QueryRow("SELECT id, login FROM users WHERE login = ? AND password = ?", login, password).Scan(&user.ID, &user.Login)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("invalid credentials")
		}
		return nil, fmt.Errorf("could not query user: %v", err)
	}

	return &user, nil
}

// Сохранение задачи в базу данных
func SaveTask(userID int, expression, result string) (*Task, error) {
	stmt, err := db.DB.Prepare("INSERT INTO tasks (expression, result, user_id) VALUES (?, ?, ?)")
	if err != nil {
		return nil, fmt.Errorf("could not prepare statement: %v", err)
	}

	res, err := stmt.Exec(expression, result, userID)
	if err != nil {
		return nil, fmt.Errorf("could not insert task: %v", err)
	}

	taskID, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("could not get last insert id: %v", err)
	}

	return &Task{ID: int(taskID), Expression: expression, Result: result, UserID: userID}, nil
}

// Получение всех задач пользователя
func GetTasksByUserID(userID int) ([]Task, error) {
	rows, err := db.DB.Query("SELECT id, expression, result FROM tasks WHERE user_id = ?", userID)
	if err != nil {
		return nil, fmt.Errorf("could not query tasks: %v", err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Expression, &task.Result); err != nil {
			return nil, fmt.Errorf("could not scan task: %v", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
