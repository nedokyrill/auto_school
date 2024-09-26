package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	// Загрузка переменных окружения из файла .env
	var DB *sql.DB
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	// Получение данных для подключения из переменных окружения
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Формирование строки подключения
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	// Подключение к базе данных
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return DB, nil
}
