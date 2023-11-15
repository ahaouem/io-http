package users

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

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
