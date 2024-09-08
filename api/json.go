package common

import (
	"encoding/json"
	"net/http"
)

func WriteJson(write http.ResponseWriter, status int, data any) {
	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(status)
	json.NewEncoder(write).Encode(data)
}

func ReadJson(read http.Request, data any) error {
	return json.NewDecoder(read.Body).Decode(data)
}

func WriteError(write http.ResponseWriter, status int, message string) {
	WriteJson(write, status, map[string]string{"error": message})
}
