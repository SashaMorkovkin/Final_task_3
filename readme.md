# Final Task 3 — REST API калькулятор на Go

## 📌 Описание

Это простой REST API на Go, реализующий калькулятор с авторизацией и хранением истории вычислений. Пользователь может зарегистрироваться, войти в систему, отправлять математические выражения на вычисление и просматривать историю операций. Все данные сохраняются в базе данных SQLite.

## 🚀 Запуск проекта

1. Убедитесь, что у вас установлен Go 1.18 или выше.
2. Склонируйте репозиторий:
   ```bash
   git clone https://github.com/SashaMorkovkin/Final_task_3.git
   cd Final_task_3


+ Запуск сервера:
    go run cmd/main.go
   > PS: Сервер работает на порте 8080

## Примеры запросов
+ Пример №1 (отправка примера)
    + Команда:
        curl --location 'localhost:8080/api/v1/calculate' \
         --header 'Authorization: Bearer your_jwt_token_here' \
         --header 'Content-Type: application/json' \
         --data '{
           "expression": "2+2*2"
        }'


    + Ответ:
        >{"id": 1, "expression": "2+2*2", "result": "6.000000"}

+ + Если калькулятор не может посчитать :
    >{"Error calculating expression"}
+ В других случаях :
    >{"Expression not found"}

+ Пример запроса для регистрации с помощью curl:
  + Команда:
       curl --location 'localhost:8080/api/v1/register' \
      --header 'Content-Type: application/json' \
      --data '{
        "login": "username",
        "password": "password123"
      }'

  + Ответ:
       {
        "id": 1,
        "login": "username"
      }

+ Пример запроса для входа с помощью curl:
   + Команда:
          curl --location 'localhost:8080/api/v1/login' \
         --header 'Content-Type: application/json' \
         --data '{
           "login": "username",
           "password": "password123"
         }'
  + Ответ:
          {
           "token": "your_jwt_token_here"
         }

+ Пример, когда пользователь ещё не авторизовался:
  + Вывод:
       > Unauthorized
  




