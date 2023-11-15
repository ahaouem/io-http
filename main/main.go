package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"server/handlers"

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

	handlers.Conn = conn

	r := chi.NewRouter()
	r.Get("/users", handlers.GetUsers)
	r.Get("/users/{id}", handlers.GetUserByID)
	r.Post("/users", handlers.CreateUser)
	r.Patch("/users/{id}", handlers.UpdateUser)
	r.Put("/users/{id}", handlers.CreateUser)
	r.Delete("/users/{id}", handlers.DeleteUser)

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
