package users

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUser(t *testing.T) {
    testCases := []struct {
        name               string
        id                 string
        expectedStatusCode int
    }{
        {
            name: "Valid ID",
            id:   "1",
            expectedStatusCode: http.StatusOK,
        },
        {
            name: "Invalid ID",
            id:   "abc",
            expectedStatusCode: http.StatusBadRequest,
        },
        {
            name: "Nonexistent ID",
            id:   "999",
            expectedStatusCode: http.StatusNotFound,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            

            req, _ := http.NewRequest("DELETE", "/users/"+tc.id, nil)
            rr := httptest.NewRecorder()

            r := chi.NewRouter()
            r.Delete("/users/{id}", DeleteUser)
            r.ServeHTTP(rr, req)

            assert.Equal(t, tc.expectedStatusCode, rr.Code, "handler returned wrong status code")
        })
    }
}