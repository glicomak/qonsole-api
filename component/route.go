package component

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func NewRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", CreateComponent(db))

	return r
}
