package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	FirstName string `json:"fname" validate:"required,alpha,min=2,max=100"`
	LastName  string `json:"lname" validate:"required,alpha,min=2,max=100"`
	Email     string `json:"email" validate:"required,unique=email,email"`
	Password  string `json:"password" validate:"required,min=8,password"`
}

func GetUserData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userData, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	fmt.Println(string(userData))

	var user User
	err = json.Unmarshal(userData, &user)
	if err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusInternalServerError)
	}
	fmt.Println(user)

	w.Write([]byte("Ok!"))
}
