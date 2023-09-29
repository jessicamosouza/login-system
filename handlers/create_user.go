package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/mail"
	"unicode"
)

type createUserPayload struct {
	FirstName string `json:"fname" validate:"required,alpha,min=2,max=100"`
	LastName  string `json:"lname" validate:"required,alpha,min=2,max=100"`
	Email     string `json:"email" validate:"required,unique=email,email"`
	Password  string `json:"password" validate:"required,min=8,password"`
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
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

	var user createUserPayload
	err = json.Unmarshal(userData, &user)
	if err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusInternalServerError)
	}
	fmt.Println(user)

	err = user.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (u createUserPayload) Validate() error {
	if err := checkName(u.FirstName); err != nil {
		return fmt.Errorf("first name validation failed: %w", err)
	}
	if err := checkName(u.LastName); err != nil {
		return fmt.Errorf("last name validation failed: %w", err)
	}
	if err := checkEmail(u.Email); err != nil {
		return fmt.Errorf("email validation failed: %w", err)
	}
	if err := checkPassword(u.Password); err != nil {
		return fmt.Errorf("password validation failed: %w", err)
	}
	return nil
}

func checkName(name string) error {
	if len(name) < 3 {
		return errors.New("name must contain at least 2 characters")
	}
	return nil
}

func checkEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("invalid email")
	}
	return nil
}

func checkPassword(password string) error {
	var (
		hasMinLen = len(password) >= 8
		hasUpper  = false
		hasLower  = false
		hasNumber = false
		hasSymbol = false
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSymbol = true
		case unicode.IsSpace(char):
			return errors.New("password cannot contain whitespace")
		}
	}

	if hasMinLen && hasUpper && hasLower && hasNumber && hasSymbol {
		return nil
	}
	return errors.New("password does not meet the requirements")
}
