package managers

import (
	"database/sql"
	"errors"
	"github.com/jessicamosouza/login-system/pkg/db"
	"github.com/jessicamosouza/login-system/pkg/models"
	"github.com/jessicamosouza/login-system/pkg/security"
	"github.com/jessicamosouza/login-system/pkg/validators"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func CreateUser(user User) error {
	err := validators.Validate(validators.User(user))
	if err != nil {
		return err
	}

	passwordHash, err := security.GenerateHash(user.Password)
	if err != nil {
		return err
	}

	err = models.CreateUser(user.FirstName, user.LastName, user.Email, passwordHash)
	if err != nil {
		return err
	}

	return nil
}

func Login(user User) error {
	err := validators.Validate(validators.User(user))
	if err != nil {
		return err
	}

	hashPassword, err := models.GetUser(user.Email, user.Password)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			return err
		}
		return err
	}

	if checkPasswordHash(user.Password, hashPassword) {
		return nil
	}

	return errors.New("invalid password")
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
