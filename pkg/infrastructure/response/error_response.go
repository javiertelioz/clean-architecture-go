package response

import (
	"encoding/json"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, status int, message string) {
	res := Response{
		Message: message,
		Data:    nil,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
	}
}
