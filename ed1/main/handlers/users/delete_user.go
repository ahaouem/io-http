package users

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    if id == "" {
        http.Error(w, "ID was not provided ", http.StatusBadRequest)
        return
    }
    userID, err := strconv.Atoi(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    commandTag, err := Conn.Exec(context.Background(), "DELETE FROM users WHERE id = $1", userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if commandTag.RowsAffected() == 0 {
        http.Error(w, "No user found with the provided ID", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
