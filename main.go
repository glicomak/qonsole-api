package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"

	"qonsole-api/auth"
	"qonsole-api/chat"
	"qonsole-api/chatspace"
	"qonsole-api/component"
	"qonsole-api/entity"
	"qonsole-api/user"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Mount("/auth", auth.NewRouter())
	r.Mount("/chat", chat.NewRouter(db))
	r.Mount("/chatspace", chatspace.NewRouter(db))
	r.Mount("/component", component.NewRouter(db))
	r.Mount("/entity", entity.NewRouter(db))
	r.Mount("/user", user.NewRouter(db))
	http.ListenAndServe(":3000", r)
}
