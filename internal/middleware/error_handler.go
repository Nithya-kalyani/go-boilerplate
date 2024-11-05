package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func JSONError(w http.ResponseWriter, statusCode int, message string) {
	response := ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Could not encode error response: %v", err)
	}
}
