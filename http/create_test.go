package http

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type ErrorReader struct{}

func (e *ErrorReader) Read([]byte) (n int, err error) {
	return 0, errors.New("simulated read error")
}

func TestCreateUser(t *testing.T) {
	t.Run("Only POST method allowed", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		rec := httptest.NewRecorder()
		CreateUserHandler(rec, req)
		if rec.Code != http.StatusMethodNotAllowed {
			t.Errorf("expected status %d but got %d", http.StatusMethodNotAllowed, rec.Code)
		}
	})

	t.Run("Handle error reading request body", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/", &ErrorReader{})
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		rec := httptest.NewRecorder()
		CreateUserHandler(rec, req)
		if rec.Code != http.StatusInternalServerError {
			t.Errorf("expected status %d but got %d", http.StatusInternalServerError, rec.Code)
		}
	})

	t.Run("Successful body read", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/",
			bytes.NewBufferString(`{"fname":"John","lname":"Doe","email":"john@email.com","password":"Password123!"}`))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		rec := httptest.NewRecorder()
		CreateUserHandler(rec, req)
		if rec.Code != http.StatusOK {
			t.Errorf("expected status %d but got %d", http.StatusOK, rec.Code)
		}

		resp := rec.Body.String()
		if resp != "LoginUserPayload created successfully!" {
			t.Errorf("expected response to be 'LoginUserPayload created successfully!' but got %s", resp)
		}
	})

}

func TestEdgeCases(t *testing.T) {
	t.Run("Invalid email", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/",
			bytes.NewBufferString(`{"fname":"John","lname":"Doe","email":"john","password":"Password123!"}`))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		rec := httptest.NewRecorder()
		CreateUserHandler(rec, req)
		if strings.TrimSpace(rec.Body.String()) != "email validation failed: invalid email" {
			t.Errorf("unexpected error message: got %v want %v", rec.Body.String(), "email validation failed: invalid email")
		}
	})

	t.Run("Invalid password", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/",
			bytes.NewBufferString(`{"fname":"John","lname":"Doe","email":"john@doe.com","password":"pass"}`))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		rec := httptest.NewRecorder()
		CreateUserHandler(rec, req)
		if strings.TrimSpace(rec.Body.String()) != "password validation failed: password does not meet the requirements" {
			t.Errorf("unexpected error message: got %v want %v", rec.Body.String(), "password validation failed: password does not meet the requirements")
		}
	})

	t.Run("Missing fields", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/",
			bytes.NewBufferString(`{"fname":"John","lname":"Doe","password":"Password123!"}`))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		rec := httptest.NewRecorder()
		CreateUserHandler(rec, req)
		if strings.TrimSpace(rec.Body.String()) != "email validation failed: invalid email" {
			t.Errorf("unexpected error message: got %v want %v", rec.Body.String(), "email validation failed: invalid email")
		}
	})

	t.Run("Invalid first name", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/",
			bytes.NewBufferString(`{"fname":"J","lname":"Doe","email":"john@doe.com", "password":"Password123!"}`))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		rec := httptest.NewRecorder()
		CreateUserHandler(rec, req)
		if strings.TrimSpace(rec.Body.String()) != "first name validation failed: name must contain at least 2 characters" {
			t.Errorf("unexpected error message: got %v want %v", rec.Body.String(), "first name validation failed: name must contain at least 2 characters")
		}
	})
}

func TestUnmarshalUser(t *testing.T) {
	t.Run("Handle error unmarshalling request body", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBufferString("Hello"))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		rec := httptest.NewRecorder()
		CreateUserHandler(rec, req)
		if rec.Code != http.StatusInternalServerError {
			t.Errorf("expected status %d but got %d", http.StatusInternalServerError, rec.Code)
		}
	})

	t.Run("Successful unmarshalling", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/",
			bytes.NewBufferString(`{"fname":"John","lname":"Doe","email":"john@email.com","password":"Password123!"}`))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		rec := httptest.NewRecorder()
		CreateUserHandler(rec, req)
		if rec.Code != http.StatusOK {
			t.Errorf("expected status %d but got %d", http.StatusOK, rec.Code)
		}
	})
}
