package chatspace

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Chatspace struct {
	Id       string `json:"id"`
	EntityId string `json:"entityId"`
}

type ChatspaceMember struct {
	ChatspaceId string `json:"chatspaceId"`
	MemberId    string `json:"memberId"`
	IsAdmin     bool   `json:"isAdmin"`
}

func CreateChatspace(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var chatspace Chatspace
		err := json.NewDecoder(r.Body).Decode(&chatspace)

		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
		}

		query := "INSERT INTO chatspace (entityid) VALUES ($1) RETURNING id"
		err = db.QueryRow(query, chatspace.EntityId).Scan(&chatspace.Id)
		if err != nil {
			http.Error(w, "Unable to create chatspace", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(chatspace)
	}
}

func AddMember(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var chatspaceMember ChatspaceMember
		err := json.NewDecoder(r.Body).Decode(&chatspaceMember)

		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
		}

		query := "INSERT INTO chatspacemember (chatspaceId, memberId) VALUES ($1, $2)"
		db.QueryRow(query, chatspaceMember.ChatspaceId, chatspaceMember.MemberId)
		chatspaceMember.IsAdmin = false

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(chatspaceMember)
	}
}
