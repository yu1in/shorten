package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type RespJson struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResp(w http.ResponseWriter, status int, message string, data interface{}) {
	resp := RespJson{
		Status:  status,
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}

func ErrorResp(w http.ResponseWriter, status int, error string) {
	resp := RespJson{
		Status: status,
		Error:  error,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}
