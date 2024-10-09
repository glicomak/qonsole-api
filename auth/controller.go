package auth

import (
	"encoding/json"
	"net/http"
)

type AuthResponse struct {
	Message string `json:"message"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	response := AuthResponse{Message: "Hello World!"}

	json.NewEncoder(w).Encode(response)
}
