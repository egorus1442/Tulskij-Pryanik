package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID модели из URL
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok || id == "" {
		http.Error(w, "ID не указан", http.StatusBadRequest)
		return
	}

	// Путь до файла с результатами
	resultPath := filepath.Join("../internal/storage/results", id+".txt")

	// Создаем директорию, если ее нет
	err := os.MkdirAll(filepath.Dir(resultPath), os.ModePerm)
	if err != nil {
		http.Error(w, "Ошибка создания директории", http.StatusInternalServerError)
		return
	}

	// Создаем файл и записываем в него заглушку
	file, err := os.Create(resultPath)
	if err != nil {
		http.Error(w, "Ошибка создания файла", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Результаты расчета")
	if err != nil {
		http.Error(w, "Ошибка записи файла", http.StatusInternalServerError)
		return
	}

	// Отправляем подтверждение клиенту
	fmt.Fprintf(w, "Расчет завершен. Результаты сохранены в %s", resultPath)
}
