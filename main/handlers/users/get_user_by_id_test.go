package users

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByID(t* testing.T) {
	testCases := []struct {
		name string
		id string
		mock func()
		expectedStatusCode int
	} {
		{
			name: "Valid ID",
			id: "1",
			expectedStatusCode: http.StatusOK,
		}, {
			name: "Invalid ID",
			id: "dupa",
			expectedStatusCode: http.StatusBadRequest,
		}, {
			name: "Nonexistent ID",
			id: "99999999",
			mock: func() {},
			expectedStatusCode: http.StatusNotFound,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()

			req, _ := http.NewRequest("GET", "/users/" + tc.id, nil)
			rr := httptest.NewRecorder()

			r := chi.NewRouter()
			r.Delete("/users/{id}", GetUserByID)
			r.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedStatusCode, rr.Code)
		})
	}	
}
