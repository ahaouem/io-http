package users

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v4"
)

func CreateUserOrUpdate(w http.ResponseWriter, r *http.Request) {
    userIDParam := chi.URLParam(r, "id")
    userID, err := strconv.Atoi(userIDParam)
    if err != nil {
        userID = 0
    }

    var userInput User
    err = json.NewDecoder(r.Body).Decode(&userInput)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if userInput.Name == "" || userInput.Lastname == "" {
        http.Error(w, "Name and lastname must be provided in the request body", http.StatusBadRequest)
        return
    }

    var user User
    err = Conn.QueryRow(context.Background(), "SELECT id FROM users WHERE id = $1", userID).Scan(&user.ID)
    if err != nil {
        if err == pgx.ErrNoRows {
            err = Conn.QueryRow(context.Background(), "INSERT INTO users (name, lastname) VALUES ($1, $2) RETURNING id", userInput.Name, userInput.Lastname).Scan(&user.ID)
            if err != nil {
                http.Error(w, "Failed to create new user: "+err.Error(), http.StatusInternalServerError)
                return
            }
            w.WriteHeader(http.StatusCreated)
            return
        } else {
            http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
            return
        }
    } else {
        _, err = Conn.Exec(context.Background(), "UPDATE users SET name = $1, lastname = $2 WHERE id = $3", userInput.Name, userInput.Lastname, userID)
        if err != nil {
            http.Error(w, "Failed to update user: "+err.Error(), http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusNoContent)
    }
}
