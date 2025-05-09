package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SashaMorkovkin/Final_task_3/internal/api"
	"github.com/SashaMorkovkin/Final_task_3/internal/calculator"
	"github.com/SashaMorkovkin/Final_task_3/internal/db"
	_ "github.com/mattn/go-sqlite3"
)

var sessionTokens = make(map[string]int)

func main() {
	// Инициализация базы данных
	db.InitDB()
	defer db.CloseDB()

	http.HandleFunc("/api/v1/register", Register)
	http.HandleFunc("/api/v1/login", Login)
	http.HandleFunc("/api/v1/calculate", Calculate)
	http.HandleFunc("/api/v1/tasks", ListTasks)

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	var creds api.User
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Регистрация пользователя
	user, err := api.RegisterUser(creds.Login, creds.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds api.User
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Авторизация пользователя
	user, err := api.AuthenticateUser(creds.Login, creds.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Генерация токена
	token := fmt.Sprintf("token_%d", user.ID)
	sessionTokens[token] = user.ID

	// Отправляем токен
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func Calculate(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	userID, ok := sessionTokens[token]
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var input struct {
		Expression string `json:"expression"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Вычисление результата
	result, err := calculator.Calculate(input.Expression)
	if err != nil {
		http.Error(w, "Calculation error: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Сохранение задачи
	task, err := api.SaveTask(userID, input.Expression, fmt.Sprintf("%f", result))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем результат
	json.NewEncoder(w).Encode(task)
}

func ListTasks(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	userID, ok := sessionTokens[token]
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	tasks, err := api.GetTasksByUserID(userID)
	if err != nil {
		http.Error(w, "Error fetching tasks", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}
