package util

import (
	"encoding/json"
	"net/http"
)

func BuildJSONResponse(w http.ResponseWriter, body any, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}
