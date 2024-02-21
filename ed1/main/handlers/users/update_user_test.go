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

func testRouter() *chi.Mux {
    r := chi.NewRouter()
    r.Put("/users/{id}", UpdateUser)
    return r
}

func TestUpdateNonExistentUser(t *testing.T) {
    router := testRouter()

    nonExistentUserID := 999
    updateData := map[string]interface{}{
        "name": "New Name",
    }
    updateDataBytes, _ := json.Marshal(updateData)
    req, _ := http.NewRequest("PUT", "/users/"+strconv.Itoa(nonExistentUserID), bytes.NewBuffer(updateDataBytes))
    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestUnitValidUserUpdate(t *testing.T) {
    validUserID := 34
    updateData := map[string]interface{}{
        "name": "Updated Name",
    }
    updateDataBytes, _ := json.Marshal(updateData)
    req, _ := http.NewRequest("PUT", "/users/"+strconv.Itoa(validUserID), bytes.NewBuffer(updateDataBytes))
    rr := httptest.NewRecorder()

    rctx := chi.NewRouteContext()
    rctx.URLParams.Add("id", strconv.Itoa(validUserID))
    req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

    UpdateUser(rr, req)

    assert.Equal(t, http.StatusNoContent, rr.Code)
    
    var user User
    err := Conn.QueryRow(context.Background(), "SELECT id, name, lastname FROM users WHERE id = $1", validUserID).Scan(&user.ID, &user.Name, &user.Lastname)
    assert.NoError(t, err)
    assert.Equal(t, "Updated Name", user.Name)
}