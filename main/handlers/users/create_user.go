package users

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    lastname := r.URL.Query().Get("lastname")
    fmt.Printf("Name: %s, Lastname: %s\n", name, lastname)

    if name == "" || lastname == "" {
        http.Error(w, "Name and lastname must be provided as query parameters", http.StatusBadRequest)
        return
    }

    user := User{
        Name:     name,
        Lastname: lastname,
    }

    
    _, err := Conn.Exec(context.Background(), "INSERT INTO users (name, lastname) VALUES ($1, $2) RETURNING id, name, lastname", user.Name, user.Lastname)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}
