package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

func PreviewHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID модели из URL через mux
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok || id == "" {
		http.Error(w, "ID не указан", http.StatusBadRequest)
		return
	}

	// Формируем путь к файлу
	filePath := filepath.Join("../internal/storage/models", id+".obj")

	// Устанавливаем заголовки для скачивания
	w.Header().Set("Content-Disposition", "attachment; filename="+id+".obj")
	w.Header().Set("Content-Type", "application/octet-stream")

	// Отправляем файл
	http.ServeFile(w, r, filePath)
}
