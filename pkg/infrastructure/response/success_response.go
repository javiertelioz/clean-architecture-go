package response

import (
	"encoding/json"
	"net/http"
)

func SuccessResponse(w http.ResponseWriter, status int, payload interface{}) {
	res := Response{
		Data:    payload,
		Message: "Operation was successful",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Failed to encode success response", http.StatusInternalServerError)
	}
}
