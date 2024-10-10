package chat

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Chat struct {
	Id          string `json:"id"`
	Text        string `json:"text"`
	SenderId    string `json:"senderId"`
	ChatspaceId string `json:"chatspaceId"`
}

func SendChat(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var chat Chat
		err := json.NewDecoder(r.Body).Decode(&chat)

		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
		}

		query := "INSERT INTO chat (text, senderid, chatspaceid) VALUES ($1, $2, $3) RETURNING id"
		err = db.QueryRow(query, chat.Text, chat.SenderId, chat.ChatspaceId).Scan(&chat.Id)
		if err != nil {
			http.Error(w, "Unable to create chat", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(chat)
	}
}
