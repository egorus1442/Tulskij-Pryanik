package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

func ApproveHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем параметр "id" из URL-пути
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "ID не указан", http.StatusBadRequest)
		return
	}

	// Проверяем, существует ли модель
	modelPath := filepath.Join("../internal/storage/models", id+".obj")
	if _, err := os.Stat(modelPath); os.IsNotExist(err) {
		http.Error(w, "Модель не найдена", http.StatusNotFound)
		return
	}

	// Подтверждаем разбиение (пока просто отправляем JSON)
	response := map[string]string{
		"message": "Разбиение модели подтверждено",
		"id":      id,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
