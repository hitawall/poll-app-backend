package handlers

import (
	"encoding/json"
	"net/http"
)

// respondJSON sends a JSON response.
func respondJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error encoding response to JSON", http.StatusInternalServerError)
	}
}
