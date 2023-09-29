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
			bytes.NewBufferString(`{"fname":"John","lname":"Doe","email":"john@email.com","password":"12345678"}`))
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
