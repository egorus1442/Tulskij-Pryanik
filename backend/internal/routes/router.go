package routes

import (
	"Tulskij-Pryanik/internal/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/upload", handlers.UploadHandler).Methods("POST")
	r.HandleFunc("/preview/{id}", handlers.PreviewHandler).Methods("GET")
	r.HandleFunc("/approve/{id}", handlers.ApproveHandler).Methods("POST")
	r.HandleFunc("/calculate/{id}", handlers.CalculateHandler).Methods("POST")
	r.HandleFunc("/status/{id}", handlers.StatusHandler).Methods("GET")
	r.HandleFunc("/result/{id}", handlers.ResultHandler).Methods("GET")

	return r
}
