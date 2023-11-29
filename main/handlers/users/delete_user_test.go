package users

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func testRouterForDeleteUser() *chi.Mux {
	r := chi.NewRouter()
	r.Delete("/users/{id}", DeleteUser)
	return r
}

func TestDeleteUserSuccess(t *testing.T) {
	router := testRouterForDeleteUser()

	validUserID := 12
	req, _ := http.NewRequest("DELETE", "/users/"+strconv.Itoa(validUserID), nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}

func TestDeleteUserInvalidID(t *testing.T) {
	router := testRouterForDeleteUser()

	req, _ := http.NewRequest("DELETE", "/users/invalidID", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDeleteUserNonExistentUser(t *testing.T) {
	router := testRouterForDeleteUser()

	nonExistentUserID := 999
	req, _ := http.NewRequest("DELETE", "/users/"+strconv.Itoa(nonExistentUserID), nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}
