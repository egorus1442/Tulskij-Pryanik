package main

import (
	"Tulskij-Pryanik/internal/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.SetupRouter()

	log.Println("Сервер запущен на :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}
