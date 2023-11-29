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

func updateRouter() *chi.Mux {
    r := chi.NewRouter()
    r.Put("/users/{id}", UpdateUser)
    return r
}

func TestUpdateUser_Success(t *testing.T) {
    router := updateRouter()
    userID := 30
    updateData := map[string]interface{}{
        "name": "Updated Name",
    }
    updateDataBytes, _ := json.Marshal(updateData)
    req, _ := http.NewRequest("PUT", "/users/"+strconv.Itoa(userID), bytes.NewBuffer(updateDataBytes))
    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)
    assert.Equal(t, http.StatusNoContent, rr.Code)
}
