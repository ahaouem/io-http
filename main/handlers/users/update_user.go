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
    if id == "" {
        http.Error(w, "ID was not provided", http.StatusBadRequest)
        return
    }
    userID, err := strconv.Atoi(id)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }
    
    var updates map[string]interface{}
    err = json.NewDecoder(r.Body).Decode(&updates)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if _, ok := updates["id"]; ok {
        http.Error(w, "Cannot update user ID", http.StatusBadRequest)
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

    commandTag, err := Conn.Exec(context.Background(), "UPDATE users SET "+query+" WHERE id = $"+strconv.Itoa(len(params)+1), append(params, userID)...)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if commandTag.RowsAffected() == 0 {
        http.Error(w, "No user found with the provided ID", http.StatusNotFound)
        return
    }

    
    user := User{
        ID: userID, 
    }
    
    for key, value := range updates {
        switch key {
        case "name":
            user.Name = value.(string)
        case "lastname":
            user.Lastname = value.(string)
        }
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
