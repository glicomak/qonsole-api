package component

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Component struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	EntityId string `json:"entityId"`
	IsRoot   bool   `json:"isRoot"`
	ParentId string `json:"parentId"`
}

func CreateComponent(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var component Component
		err := json.NewDecoder(r.Body).Decode(&component)

		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
		}

		query := "INSERT INTO component (name, entityid, isroot, parentid) VALUES ($1, $2, $3, $4) RETURNING id"
		err = db.QueryRow(query, component.Name, component.EntityId, component.IsRoot, component.ParentId).Scan(&component.Id)
		if err != nil {
			http.Error(w, "Unable to create component", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(component)
	}
}
