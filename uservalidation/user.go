package uservalidation

import (
	"errors"
	"fmt"
	"net/mail"
	"regexp"
	"unicode"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func Validate(user User, isSignUp bool) error {
	if err := CheckEmail(user.Email); err != nil {
		return fmt.Errorf("email validation failed: %w", err)
	}
	if err := checkPassword(user.Password); err != nil {
		return fmt.Errorf("password validation failed: %w", err)
	}

	if isSignUp {
		if err := CheckName(user.FirstName); err != nil {
			return fmt.Errorf("first name validation failed: %w", err)
		}
		if err := CheckName(user.LastName); err != nil {
			return fmt.Errorf("last name validation failed: %w", err)
		}
	}

	return nil
}

func CheckName(name string) error {
	re := regexp.MustCompile(`[^a-zA-ZäöüÄÖÜáéíóúÁÉÍÓÚàèìòùÀÈÌÒÙâêîôûÂÊÎÔÛãñõÃÑÕçÇ ^\p{L}'-]`)

	if name == "" {
		return errors.New("name cannot be empty")
	} else if len(name) < 2 {
		return errors.New("name must contain at least 2 characters")
	} else if re.MatchString(name) {
		return errors.New("name cannot contain special characters")
	}

	return nil
}

func CheckEmail(email string) error {
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
