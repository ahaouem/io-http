package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

func main() {
    wd, err := os.Getwd()
    if err != nil {
        log.Fatal("Error getting working directory: ", err)
    }
    fmt.Println("Current working directory:", wd)

    err = godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    databaseURL := os.Getenv("DATABASE_URL")
    if databaseURL == "" {
        fmt.Println("DATABASE_URL is not set")
        return
    }

    conn, err := pgx.Connect(context.Background(), databaseURL)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
        os.Exit(1)
    }
    defer conn.Close(context.Background())

    fmt.Println("Successfully connected to database")

    rows, err := conn.Query(context.Background(), `SELECT * FROM users`)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to query users table: %v\n", err)
        os.Exit(1)
    }
    defer rows.Close()

    for rows.Next() {
        var (
            id       int
            name     string
            lastname string
        )

        err = rows.Scan(&id, &name, &lastname)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Unable to read row: %v\n", err)
            continue
        }
        fmt.Printf("ID: %d, Name: %s, Lastname: %s\n", id, name, lastname)
    }

    if err = rows.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "Error iterating through rows: %v\n", err)
    }
}
