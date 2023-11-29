package http

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginUserHandler(t *testing.T) {
	t.Parallel()

	t.Run("Errors", testLoginErrors)
	t.Run("Success", testLoginSuccess)
}

func testLoginErrors(t *testing.T) {
	t.Parallel()

	methodsToTest := []string{http.MethodDelete, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodOptions}

	for _, method := range methodsToTest {
		t.Run(fmt.Sprintf("Method %s not allowed", method), func(t *testing.T) {
			t.Parallel()
			req, err := http.NewRequest(method, "/", nil)
			require.NoError(t, err)

			resp := assertResponseStatus(t, req, http.StatusMovedPermanently)
			require.Equal(t, "", resp)
		})
	}

	t.Run("Handle error reading request body", func(t *testing.T) {
		t.Parallel()
		req, err := http.NewRequest(http.MethodGet, "/", &ErrorReader{})
		require.NoError(t, err)

		resp := assertResponseStatus(t, req, http.StatusInternalServerError)
		require.Equal(t, "Error unmarshalling request body\nemail validation failed: invalid email\n", resp)
	})

	t.Run("Empty body", func(t *testing.T) {
		t.Parallel()
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		require.NoError(t, err)

		resp := assertResponseStatus(t, req, http.StatusBadRequest)
		require.Equal(t, "Empty body\n", resp)
	})

	t.Run("Error finding user", func(t *testing.T) {
		t.Parallel()
		req, err := http.NewRequest(http.MethodGet, "/", bytes.NewBufferString(
			`{"email":"maria@email.com", "password":"Password123!"}`))
		require.NoError(t, err)

		resp := assertResponseStatus(t, req, http.StatusBadRequest)
		require.Equal(t, "[usermodels] user not found\n", resp)
	})
}

func testLoginSuccess(t *testing.T) {
	t.Parallel()

	t.Run("Successful body read", func(t *testing.T) {
		t.Parallel()
		req, err := http.NewRequest(http.MethodGet, "/",
			bytes.NewBufferString(`{"email":"john@doe.com","password":"Password123!"}`))
		require.NoError(t, err)

		resp := assertResponseStatus(t, req, http.StatusOK)
		require.Equal(t, "LoginUserPayload logged successfully!", resp)
	})
}

func assertResponseStatus(t *testing.T, req *http.Request, status int) string {
	rec := httptest.NewRecorder()
	LoginUserHandler(rec, req)
	require.Equal(t, status, rec.Code)
	return rec.Body.String()
}
