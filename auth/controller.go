package auth

import (
	"net/http"

	"encoding/json"
)

type AuthResponse struct {
	Message string `json:"message"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	response := AuthResponse{Message: "Hello World!"}

	json.NewEncoder(w).Encode(response)
}