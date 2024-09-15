package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"qonsole-api/auth"
)

func main() {
	r := chi.NewRouter()

	r.Mount("/auth", auth.NewRouter())

	http.ListenAndServe(":3000", r)
}