package users

import (
	"context"
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

func TestGetUserByIDInvalidID(t *testing.T) {
	req, _ := http.NewRequest("GET", "/users/invalidID", nil)
	rr := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "invalidID")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	GetUserByID(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetUserByIDNonExistentUser(t *testing.T) {
	nonExistentUserID := 999
	req, _ := http.NewRequest("GET", "/users/"+strconv.Itoa(nonExistentUserID), nil)
	rr := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", strconv.Itoa(nonExistentUserID))
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	GetUserByID(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}