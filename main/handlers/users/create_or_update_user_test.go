package users

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func createUserOrUpdateRouter() *chi.Mux {
    r := chi.NewRouter()
    r.Post("/users/{id}", CreateUserOrUpdate)
    return r
}

func TestCreateUserOrUpdate_Create(t *testing.T) {
    router := createUserOrUpdateRouter()
    newUser := User{
        Name:     "New User",
        Lastname: "Test",
    }
    newUserBytes, _ := json.Marshal(newUser)
    req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(newUserBytes))
    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)
    assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestCreateUserOrUpdate_Update(t *testing.T) {
    router := createUserOrUpdateRouter()
    userID := 41
    updatedUser := User{
        Name:     "Updated User",
        Lastname: "Test",
    }
    updatedUserBytes, _ := json.Marshal(updatedUser)
    req, _ := http.NewRequest("POST", "/users/"+strconv.Itoa(userID), bytes.NewBuffer(updatedUserBytes))
    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)
    assert.Equal(t, http.StatusNoContent, rr.Code)
}
