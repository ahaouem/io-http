package users

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)


func TestUpdateUserTest(t *testing.T) {
	testCases := []struct {
		name string
		id string
		RequestName string
		lastname string
		mock func()
		expectedStatusCode int
	} {
		{
			name: "Valid ID",
			id: "1",
			RequestName: "Aleksander",
			lastname: "Kowalski",
			expectedStatusCode: http.StatusOK,
		}, {
			name: "Blank name field",
			id: "1",
			RequestName: "",
			lastname: "Kowalski",
			expectedStatusCode: http.StatusBadRequest,
		}, {
			name: "Blank lastname field",
			id: "5",
			RequestName: "Aleksander",
			lastname: "",
			expectedStatusCode: http.StatusBadRequest,
		}, {
			name: "Invalid ID",
			id: "dupa",
			RequestName: "Aleksander",
			lastname: "Kowalski",
			expectedStatusCode: http.StatusBadRequest,
		}, {
			name: "Nonexistent ID",
			id: "99999999",
			RequestName: "Aleksander",
			lastname: "Kowalski",
			mock: func() {},
			expectedStatusCode: http.StatusNotFound,
		},
		}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()

			req, _ := http.NewRequest("PUT", "/users/" + tc.id, nil)
			rr := httptest.NewRecorder()

			r := chi.NewRouter()
			r.Put("/users/{id}", UpdateUser)
			r.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedStatusCode, rr.Code)
		})
	}
}