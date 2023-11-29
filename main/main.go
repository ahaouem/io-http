package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"server/handlers/users"

	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	users.Conn = conn

	r := chi.NewRouter()
	r.Get("/users", users.GetUsers)
	r.Get("/users/{id}", users.GetUserByID)
	r.Post("/users", users.CreateUserOrUpdate)
	r.Patch("/users/{id}", users.UpdateUser)
	r.Put("/users/{id}", users.CreateUserOrUpdate)
	r.Delete("/users/{id}", users.DeleteUser)

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
