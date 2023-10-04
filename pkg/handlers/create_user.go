package handlers

import (
	"encoding/json"
	"github.com/jessicamosouza/login-system/pkg/managers"
	"net/http"
)

type CreateUserPayload struct {
	FirstName string `json:"fname" validate:"required,alpha,min=2,max=20"`
	LastName  string `json:"lname" validate:"required,alpha,min=2,max=100"`
	Email     string `json:"email" validate:"required,unique=email,email"`
	Password  string `json:"password" validate:"required,min=8,password"`
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user CreateUserPayload
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusInternalServerError)
	}

	err := managers.CreateUser(managers.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User created successfully!"))
}
