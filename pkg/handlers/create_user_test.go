package handlers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type errorReader struct{}

func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("simulated read error")
}

func TestGetUser(t *testing.T) {
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
		req, err := http.NewRequest(http.MethodPost, "/", &errorReader{})
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
		req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBufferString("Hello"))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		rec := httptest.NewRecorder()
		CreateUserHandler(rec, req)
		if rec.Code != http.StatusOK {
			t.Errorf("expected status %d but got %d", http.StatusOK, rec.Code)
		}
		resp := rec.Body.String()
		if resp != "Ok!" {
			t.Errorf("expected response to be 'Ok!' but got %s", resp)
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

func TestCreateUserPayload_Validate(t *testing.T) {
	t.Run("First name validation failed", func(t *testing.T) {
		u := CreateUserPayload{
			FirstName: "123",
			LastName:  "Doe",
			Email:     "doe@email.com",
			Password:  "Password123!",
		}
		err := u.Validate()
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("Last name validation failed", func(t *testing.T) {
		u := CreateUserPayload{
			FirstName: "John",
			LastName:  "123",
			Email:     "doe@email.com",
			Password:  "Password123!",
		}
		err := u.Validate()
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("Email validation failed", func(t *testing.T) {
		u := CreateUserPayload{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "doe",
			Password:  "Password123!",
		}
		err := u.Validate()
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("Password validation failed", func(t *testing.T) {
		u := CreateUserPayload{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "doe",
			Password:  "123",
		}
		err := u.Validate()
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("Successful validation", func(t *testing.T) {
		u := CreateUserPayload{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "doe",
			Password:  "Password123!",
		}
		err := u.Validate()
		if err != nil {
			t.Errorf("expected nil but got %v", err)
		}
	})
}

func TestCreateUserHandler(t *testing.T) {
	tests := []struct {
		FirstName string
		LastName  string
		Email     string
		Password  string
	}{
		{"123", "Doe", "doe@email.com", "Password123!"},
		{"John", "123", "doe@email.com", "Password123!"},
		{"John", "Doe", "doe", "Password123!"},
		{"John", "Doe", "doe@email.com", "123"},
		{"John", "Doe", "doe", "123"},
	}
	for _, tt := range tests {
		t.Run("First name validation failed", func(t *testing.T) {
			u := CreateUserPayload{
				FirstName: tt.FirstName,
				LastName:  tt.LastName,
				Email:     tt.Email,
				Password:  tt.Password,
			}
			err := u.Validate()
			if err == nil {
				t.Errorf("expected error but got nil")
			}
		})
	}
}
