package users

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func testRouterForGetUsers() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/users", GetUsers)
	return r
}

func TestGetUsersSuccess(t *testing.T) {
	router := testRouterForGetUsers()

	req, _ := http.NewRequest("GET", "/users", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var users []User
	err := json.Unmarshal(rr.Body.Bytes(), &users)
	assert.NoError(t, err)
}

