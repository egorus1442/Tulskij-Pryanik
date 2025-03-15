package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

func ResultHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID из URL
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok || id == "" {
		http.Error(w, "ID не указан", http.StatusBadRequest)
		return
	}

	// Путь к файлу с результатами
	resultPath := filepath.Join("../internal/storage/results", id+".txt")

	// Проверяем, существует ли файл
	if _, err := os.Stat(resultPath); os.IsNotExist(err) {
		http.Error(w, "Результаты не найдены", http.StatusNotFound)
		return
	}

	// Отправляем файл пользователю
	w.Header().Set("Content-Disposition", "attachment; filename=result.txt")
	w.Header().Set("Content-Type", "text/plain")
	http.ServeFile(w, r, resultPath)
}
