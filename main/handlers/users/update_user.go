package users

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

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
