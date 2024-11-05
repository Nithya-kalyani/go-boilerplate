package utils

import (
	"encoding/json"
	"net/http"
)

func RenderJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}
}
