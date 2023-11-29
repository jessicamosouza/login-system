package http

import (
	"encoding/json"
	"github.com/jessicamosouza/login-system/userops"
	"net/http"
)

type LoginUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

	if r.Body == nil {
		http.Error(w, "Empty body", http.StatusBadRequest)
		return
	}

	var user LoginUserPayload
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusInternalServerError)
	}

	err := userops.Login(userops.User{
		Email:    user.Email,
		Password: user.Password,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("LoginUserPayload logged successfully!"))
}
