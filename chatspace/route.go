package chatspace

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func NewRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", CreateChatspace(db))
	r.Post("/member", AddMember(db))

	return r
}
