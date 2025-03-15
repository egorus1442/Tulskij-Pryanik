package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

// Структура запроса для параметров расчета
type UploadRequest struct {
	Pressure     float64 `json:"pressure"`
	FlowSpeed    float64 `json:"flow_speed"`
	Density      float64 `json:"density"`
	PanelDensity int     `json:"panel_density"`
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Ограничение на размер файла (например, 50MB)
	r.ParseMultipartForm(50 << 20)

	// Получаем файл
	file, handler, err := r.FormFile("model")
	if err != nil {
		log.Printf("Ошибка загрузки файла: %v", err)
		http.Error(w, "Ошибка загрузки файла", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Генерируем уникальный ID
	id := uuid.New().String()

	// Сохраняем файл локально
	savePath := filepath.Join("../internal/storage/models", id+".obj")
	out, err := os.Create(savePath)
	if err != nil {
		log.Printf("Ошибка сохранения файла: %v", err)
		http.Error(w, "Ошибка сохранения файла", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Printf("Ошибка записи файла: %v", err)
		http.Error(w, "Ошибка записи файла", http.StatusInternalServerError)
		return
	}

	// Читаем параметры из запроса						//потом раскоментить
	//	params := UploadRequest{
	//		Pressure:     getFloatParam(r, "pressure"),
	//		FlowSpeed:    getFloatParam(r, "flow_speed"),
	//		Density:      getFloatParam(r, "density"),
	//		PanelDensity: int(getFloatParam(r, "panel_density")),
	//	}

	// Выводим успешный ответ
	response := map[string]string{
		"id":      id,
		"message": fmt.Sprintf("Файл %s загружен. ID: %s", handler.Filename, id),
	}
	json.NewEncoder(w).Encode(response)
}

// Вспомогательная функция для получения float параметра						//потом раскоментить
//func getFloatParam(r *http.Request, key string) float64 {
//	val := r.FormValue(key)
//	if val == "" {
//		return 0.0
//	}
//	var result float64
//	fmt.Sscanf(val, "%f", &result)
//	return result
//}
