package users

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func getUsersRouter() *chi.Mux {
    r := chi.NewRouter()
    r.Get("/users", GetUsers)
    return r
}

func TestGetUsers_Success(t *testing.T) {
    router := getUsersRouter()
    req, _ := http.NewRequest("GET", "/users", nil)
    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)
    assert.Equal(t, http.StatusOK, rr.Code)
}
