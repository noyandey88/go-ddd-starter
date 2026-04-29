package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Status  string `json:"status"`
	Payload any    `json:"payload"`
}

func SendData(w http.ResponseWriter, success bool, message string, data any, statusCode int) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(Response{
		Success: success,
		Message: message,
		Status:  http.StatusText(statusCode),
		Payload: data,
	})

	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: err.Error(),
			Status:  http.StatusText(statusCode),
			Payload: nil,
		})
	}
}

func SendError(w http.ResponseWriter, success bool, message string, data any, statusCode int) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(Response{
		Success: success,
		Message: message,
		Status:  http.StatusText(statusCode),
		Payload: data,
	})

	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: err.Error(),
			Status:  http.StatusText(http.StatusInternalServerError),
			Payload: nil,
		})
	}
}
