package users

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func getUserByIDRouter() *chi.Mux {
    r := chi.NewRouter()
    r.Get("/users/{id}", GetUserByID)
    return r
}

func TestGetUserByID_Success(t *testing.T) {
    router := getUserByIDRouter()
    userID := 30
    req, _ := http.NewRequest("GET", "/users/"+strconv.Itoa(userID), nil)
    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)
    assert.Equal(t, http.StatusOK, rr.Code)
}
