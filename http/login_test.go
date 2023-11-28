package http

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	t.Run("Only GET method allowed", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/", nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		rec := httptest.NewRecorder()
		LoginUserHandler(rec, req)
		if rec.Code != http.StatusMovedPermanently {
			t.Errorf("expected status %d but got %d", http.StatusMethodNotAllowed, rec.Code)
		}
	})

	t.Run("Handle error reading request body", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", &ErrorReader{})
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		rec := httptest.NewRecorder()
		LoginUserHandler(rec, req)
		if rec.Code != http.StatusInternalServerError {
			t.Errorf("expected status %d but got %d", http.StatusInternalServerError, rec.Code)
		}
	})

	t.Run("Successful body read", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/",
			bytes.NewBufferString(`{"email":"john@doe.com","password":"Password123!"}`))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		rec := httptest.NewRecorder()
		LoginUserHandler(rec, req)
		if rec.Code != http.StatusOK {
			t.Errorf("expected status %d but got %d", http.StatusOK, rec.Code)
		}

		resp := rec.Body.String()
		if resp != "LoginUserPayload logged successfully!" {
			t.Errorf("expected response to be 'LoginUserPayload logged successfully!' but got %s", resp)
		}
	})

	t.Run("Empty body", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		rec := httptest.NewRecorder()
		LoginUserHandler(rec, req)
		if rec.Code != http.StatusBadRequest {
			t.Errorf("expected status %d but got %d", http.StatusBadRequest, rec.Code)
		}

		resp := rec.Body.String()
		if strings.TrimSpace(resp) != "Empty body" {
			t.Errorf("expected response to be 'Empty body' but got %s", resp)
		}
	})

	t.Run("Error finding user", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", bytes.NewBufferString(`{"email":"maria@email.com",
"password":"Password123!"}`))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		rec := httptest.NewRecorder()
		LoginUserHandler(rec, req)
		if rec.Code != http.StatusBadRequest {
			t.Errorf("expected status %d but got %d", http.StatusBadRequest, rec.Code)
		}
	})
}
