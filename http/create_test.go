package http

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type ErrorReader struct{}

func (e *ErrorReader) Read([]byte) (n int, err error) {
	return 0, errors.New("simulated read error")
}

const successfulBody = `{"fname":"John","lname":"Doe","email":"john@email.com","password":"Password123!"}`

func TestCreateUser(t *testing.T) {
	t.Run("Only POST method allowed", func(t *testing.T) {
		req, rec := createRequest(t, http.MethodGet, "/", "")
		CreateUserHandler(rec, req)
		checkResponseCode(t, rec, http.StatusMethodNotAllowed)
	})

	t.Run("Handle error reading request body", func(t *testing.T) {
		errReader := io.NopCloser(bytes.NewReader(nil))
		req, _ := http.NewRequest(http.MethodPost, "/", errReader)
		rec := httptest.NewRecorder()
		CreateUserHandler(rec, req)
		checkResponseCode(t, rec, http.StatusInternalServerError)
	})

	t.Run("Successful body read", func(t *testing.T) {
		req, rec := createRequest(t, http.MethodPost, "/", successfulBody)
		CreateUserHandler(rec, req)
		checkResponseCode(t, rec, http.StatusOK)

		resp := rec.Body.String()
		if resp != "LoginUserPayload created successfully!" {
			t.Errorf("expected response to be 'LoginUserPayload created successfully!' but got %s", resp)
		}
	})

}

func TestEdgeCases(t *testing.T) {
	testCases := []struct {
		name          string
		requestBody   string
		expectedCode  int
		expectedError string
	}{
		{
			"Invalid email",
			`{"fname":"John","lname":"Doe","email":"john","password":"Password123!"}`,
			http.StatusBadRequest,
			"email validation failed: invalid email",
		},
		{
			"Invalid password",
			`{"fname":"John","lname":"Doe","email":"john@doe.com","password":"pass"}`,
			http.StatusBadRequest,
			"password validation failed: password does not meet the requirements",
		},
		{
			"Missing field",
			`{"fname":"John","lname":"Doe","password":"Password123!"}`,
			http.StatusBadRequest,
			"email validation failed: invalid email",
		},
		{
			"Missing more than one field",
			`{"fname":"John", "password":"pass"}`,
			http.StatusBadRequest,
			"email validation failed: invalid email",
		},
		{
			"Invalid first name",
			`{"fname":"J","lname":"Doe","email":"john@doe.com", "password":"Password123!"}`,
			http.StatusBadRequest,
			"first name validation failed: name must contain at least 2 characters",
		},
		{
			"Invalid last name",
			`{"fname":"John","lname":"D","email":"john@doe.com","password":"Password123!"}`,
			http.StatusBadRequest,
			"last name validation failed: name must contain at least 2 characters",
		},
		{
			"Invalid password",
			`{"fname":"John","lname":"Doe","email":"john@doe.com","password":"pass"}`,
			http.StatusBadRequest,
			"password validation failed: password does not meet the requirements",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, rec := createRequest(t, http.MethodPost, "/", tc.requestBody)
			CreateUserHandler(rec, req)
			checkResponseCode(t, rec, tc.expectedCode)
			checkResponseMessage(t, rec, tc.expectedError)
		})
	}
}

func TestUnmarshalUser(t *testing.T) {
	t.Run("Handle error unmarshalling request body", func(t *testing.T) {
		req, rec := createRequest(t, http.MethodPost, "/", "Hello")
		CreateUserHandler(rec, req)
		checkResponseCode(t, rec, http.StatusInternalServerError)
	})

	t.Run("Successful unmarshalling", func(t *testing.T) {
		req, rec := createRequest(t, http.MethodPost, "/", successfulBody)
		CreateUserHandler(rec, req)
		checkResponseCode(t, rec, http.StatusOK)
	})
}

func createRequest(t *testing.T, method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req, err := http.NewRequest(method, url, bytes.NewBufferString(body))
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	return req, httptest.NewRecorder()
}

func checkResponseCode(t *testing.T, rec *httptest.ResponseRecorder, expectedCode int) {
	if rec.Code != expectedCode {
		t.Errorf("expected status %d but got %d", expectedCode, rec.Code)
	}
}

func checkResponseMessage(t *testing.T, rec *httptest.ResponseRecorder, expectedMessage string) {
	if strings.TrimSpace(rec.Body.String()) != expectedMessage {
		t.Errorf("unexpected error message: got %v want %v", rec.Body.String(), expectedMessage)
	}
}
