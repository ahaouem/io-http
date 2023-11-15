package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v4"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

var Conn *pgx.Conn

func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := Conn.Query(context.Background(), `SELECT id, name, lastname FROM users`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Lastname); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userID, _ := strconv.Atoi(id)
	var user User
	err := Conn.QueryRow(context.Background(), "SELECT id, name, lastname FROM users WHERE id = $1", userID).
			Scan(&user.ID, &user.Name, &user.Lastname)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

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


func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userID, _ := strconv.Atoi(id)
	var updates map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
	}
	var query string
	var params []interface{}
	for key, value := range updates {
			query += key + " = $" + strconv.Itoa(len(params)+1) + ", "
			params = append(params, value)
	}
	if query == "" {
			http.Error(w, "No valid fields to update", http.StatusBadRequest)
			return
	}
	query = query[:len(query)-2]
	_, err = Conn.Exec(context.Background(), "UPDATE users SET "+query+" WHERE id = $"+strconv.Itoa(len(params)+1)+" RETURNING id, name, lastname", append(params, userID)...)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}
	var user User
	user.ID = userID
	for key, value := range updates {
			if key == "name" {
					user.Name = value.(string)
			} else if key == "lastname" {
					user.Lastname = value.(string)
			}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(id)
	if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
	}
	_, err = Conn.Exec(context.Background(), "DELETE FROM users WHERE id = $1", userID)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}
	w.WriteHeader(http.StatusOK)
}