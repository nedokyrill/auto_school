package main

import (
	"database/sql"
	"log"
	"newWebServer/cmd/api"
	"newWebServer/db"
)

func main() {
	//Подключение к бд
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	initStorage(db)

	//Подключение к серверу
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

// Проверка подключения
func initStorage(DB *sql.DB) {
	err := DB.Ping()
	if err != nil {
		log.Fatal("Ошибка при попытке соединения с базой данных: ", err)
	}
	log.Println("Успешное подключение к базе данных!")
}
