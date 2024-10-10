package user

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	EntityId  string `json:"entityId"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
		}

		query := "INSERT INTO \"user\" (name, entityid, password, firstname, lastname) VALUES ($1, $2, $3, $4, $5) RETURNING id"
		err = db.QueryRow(query, user.Name, user.EntityId, user.Password, user.FirstName, user.LastName).Scan(&user.Id)
		if err != nil {
			http.Error(w, "Unable to create user", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}
