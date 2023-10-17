package handlers

import (
	"encoding/json"
	"github.com/jessicamosouza/login-system/pkg/managers"
	"net/http"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

	if r.Body == nil {
		http.Error(w, "Empty body", http.StatusBadRequest)
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusInternalServerError)
	}

	err := managers.Login(managers.User{
		Email:    user.Email,
		Password: user.Password,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User logged successfully!"))
}
