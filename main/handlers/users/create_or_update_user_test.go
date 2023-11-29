package users

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateUserTest(t *testing.T) {
    testCases := []struct {
        name               string
        input              map[string]string
        expected           User
        expectedStatusCode int
    }{
        {
            name:     "Valid name and lastname",
            input:    map[string]string{"name": "John", "lastname": "Doe"},
            expected: User{Name: "John", Lastname: "Doe"},
            expectedStatusCode: http.StatusOK,
        },
        {
            name:     "Blank name",
            input:    map[string]string{"name": "", "lastname": "Doe"},
            expected: User{Name: "", Lastname: "Doe"},
            expectedStatusCode: http.StatusBadRequest,
        },
        {
            name:     "Blank lastname",
            input:    map[string]string{"name": "John", "lastname": ""},
            expected: User{Name: "John", Lastname: ""},
            expectedStatusCode: http.StatusBadRequest,
        },
        {
            name:     "Both name and lastname blank",
            input:    map[string]string{"name": "", "lastname": ""},
            expected: User{Name: "", Lastname: ""},
                    expectedStatusCode: http.StatusBadRequest,
        },
        {
            name:     "Name field missing",
            input:    map[string]string{"lastname": "Doe"},
            expected: User{Name: "", Lastname: "Doe"},
                    expectedStatusCode: http.StatusBadRequest,
        },
        {
            name:     "Lastname field missing",
            input:    map[string]string{"name": "John"},
            expected: User{Name: "John", Lastname: ""},
                    expectedStatusCode: http.StatusBadRequest,
        },
        {
            name:     "Both fields missing",
            input:    map[string]string{},
            expected: User{Name: "", Lastname: ""},
                    expectedStatusCode: http.StatusBadRequest,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            userDataBytes, _ := json.Marshal(tc.input)
            req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(userDataBytes))
            rr := httptest.NewRecorder()

            CreateUserOrUpdate(rr, req)

            assert.Equal(t, tc.expectedStatusCode, rr.Code, "handler returned wrong status code")

            var user User
            _ = json.Unmarshal(rr.Body.Bytes(), &user)

            assert.Equal(t, tc.expected.Name, user.Name, "Expected Name did not match")
            assert.Equal(t, tc.expected.Lastname, user.Lastname, "Expected Lastname did not match")
        })
    }
}