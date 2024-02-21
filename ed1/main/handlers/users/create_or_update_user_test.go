package users

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func testRouterForCreateOrUpdateUser() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/users/{id}", CreateUserOrUpdate)
	r.Put("/users/{id}", CreateUserOrUpdate)
	return r
}

func TestCreateUserWithPost(t *testing.T) {
	router := testRouterForCreateOrUpdateUser()

	userInput := User{Name: "New", Lastname: "User"}
	userInputBytes, _ := json.Marshal(userInput)
	req, _ := http.NewRequest("POST", "/users/60", bytes.NewBuffer(userInputBytes)) 
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestUpdateUserWithPut(t *testing.T) {
	router := testRouterForCreateOrUpdateUser()

	validUserID := 42 
	userInput := User{Name: "Updated", Lastname: "User"}
	userInputBytes, _ := json.Marshal(userInput)
	req, _ := http.NewRequest("PUT", "/users/"+strconv.Itoa(validUserID), bytes.NewBuffer(userInputBytes))
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}


func TestBadUpdateUserWithPut(t *testing.T) {
	router := testRouterForCreateOrUpdateUser()

	validUserID := 1
	userInput := User{Name: "Updated", Lastname: "User"}
	userInputBytes, _ := json.Marshal(userInput)
	req, _ := http.NewRequest("PUT", "/users/"+strconv.Itoa(validUserID), bytes.NewBuffer(userInputBytes))
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}


func TestUnitCreateUserOrUpdateCreate(t *testing.T) {
	userInput := User{Name: "New", Lastname: "User"}
	userInputBytes, _ := json.Marshal(userInput)
	req, _ := http.NewRequest("POST", "/users/0", bytes.NewBuffer(userInputBytes))
	rr := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "0")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	CreateUserOrUpdate(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestUnitCreateUserOrUpdateUpdate(t *testing.T) {
	validUserID := 42
	userInput := User{Name: "Updated", Lastname: "User"}
	userInputBytes, _ := json.Marshal(userInput)
	req, _ := http.NewRequest("PUT", "/users/"+strconv.Itoa(validUserID), bytes.NewBuffer(userInputBytes))
	rr := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", strconv.Itoa(validUserID))
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))


	CreateUserOrUpdate(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}