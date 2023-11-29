package users

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func testRouterForGetUserByID() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/users/{id}", GetUserByID)
	return r
}

func TestGetUserByIDSuccess(t *testing.T) {
	router := testRouterForGetUserByID()

	validUserID := 41
	req, _ := http.NewRequest("GET", "/users/"+strconv.Itoa(validUserID), nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var user User
	err := json.Unmarshal(rr.Body.Bytes(), &user)
	assert.NoError(t, err)
	assert.Equal(t, validUserID, user.ID)
	
}

func TestGetUserByIDInvalidID(t *testing.T) {
	router := testRouterForGetUserByID()

	req, _ := http.NewRequest("GET", "/users/invalidID", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetUserByIDNonExistentUser(t *testing.T) {
	router := testRouterForGetUserByID()
	nonExistentUserID := 999
	req, _ := http.NewRequest("GET", "/users/"+strconv.Itoa(nonExistentUserID), nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
