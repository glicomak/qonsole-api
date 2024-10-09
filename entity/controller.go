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

type User struct {
	Id			string	`json:"id"`
	Name		string	`json:"name"`
	EntityId	string	`json:"entityId"`
	Password	string	`json:"password"`
	FirstName	string	`json:"firstName"`
	LastName	string	`json:"lastName"`
}

type EntityUserPair struct {
	Entity	Entity	`json:"entity"`
	User	User	`json:"user"`
}

func CreateEntity(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request EntityUserPair
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		entity := request.Entity
		user := request.User

		query := "INSERT INTO entity (name) VALUES ($1) RETURNING id"
		err = db.QueryRow(query, entity.Name).Scan(&entity.Id)
		if err != nil {
			http.Error(w, "Unable to create entity", http.StatusInternalServerError)
			return
		}
		user.EntityId = entity.Id

		query = "INSERT INTO \"user\" (name, entityid, password, firstname, lastname) VALUES ($1, $2, $3, $4, $5) RETURNING id"
		err = db.QueryRow(query, user.Name, user.EntityId, user.Password, user.FirstName, user.LastName).Scan(&user.Id)
		if err != nil {
			http.Error(w, "Unable to create user", http.StatusInternalServerError)
			return
		}

		response := EntityUserPair{
			Entity: entity,
			User: user,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}
