package handlers

import (
	"ShishaSabourAPI/internal/logger"
	"ShishaSabourAPI/internal/models"
	"ShishaSabourAPI/internal/service"
	"encoding/json"
	"net/http"
)

type tobbacoHandler struct{}

func NewTobbacoHandler() TobbacoHandler {
	return &tobbacoHandler{}
}

func (h *tobbacoHandler) GetTobbacos(w http.ResponseWriter, r *http.Request) {
	var tobaccos []models.Tobacco

	tobaccos = service.GetTobbacosDB()

	jsonData, err := json.Marshal(tobaccos)
	if err != nil {
		http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
		logger.Errorf("Error convert to JSON, Error: %v", http.StatusInternalServerError)
		return
	}

	if err != nil {
		http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
		logger.Errorf("Error convert to JSON, Error: %v", http.StatusInternalServerError)
		return
	}

	// Establecer la cabecera Content-Type a "application/json"
	w.Header().Set("Content-Type", "application/json")

	// Escribir la respuesta JSON
	w.Write(jsonData)
}
