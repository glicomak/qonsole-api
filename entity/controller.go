package entity

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Entity struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func createEntity(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var entity Entity
		err := json.NewDecoder(r.Body).Decode(&entity)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		query := "INSERT INTO entity (name) VALUES ($1) RETURNING id"
		err = db.QueryRow(query, entity.Name).Scan(&entity.Id)
		if err != nil {
			http.Error(w, "Unable to create entity", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(entity)
	}
}
