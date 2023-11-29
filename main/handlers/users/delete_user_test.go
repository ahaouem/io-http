package users

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func deleteUserRouter() *chi.Mux {
    r := chi.NewRouter()
    r.Delete("/users/{id}", DeleteUser)
    return r
}

func TestDeleteUser_Success(t *testing.T) {
    router := deleteUserRouter()
    userID := 40
    req, _ := http.NewRequest("DELETE", "/users/"+strconv.Itoa(userID), nil)
    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)
    assert.Equal(t, http.StatusNoContent, rr.Code)
}
