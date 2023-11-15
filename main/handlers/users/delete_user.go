package users

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

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
    w.Write([]byte("User deleted"))
}
