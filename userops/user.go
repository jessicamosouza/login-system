package userops

import (
	"errors"
	"github.com/jessicamosouza/login-system/security"
	"github.com/jessicamosouza/login-system/usermodels"
	"github.com/jessicamosouza/login-system/uservalidation"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func CreateUser(user User) error {
	err := uservalidation.Validate(uservalidation.User(user), true)
	if err != nil {
		return err
	}

	passwordHash, err := security.GenerateHash(user.Password)
	if err != nil {
		return err
	}

	err = usermodels.CreateUser(user.FirstName, user.LastName, user.Email, passwordHash)
	if err != nil {
		return err
	}

	return nil
}

func Login(user User) error {
	err := uservalidation.Validate(uservalidation.User(user), false)
	if err != nil {
		return err
	}

	hashPassword, err := usermodels.GetUser(user.Email, user.Password)
	if err != nil {
		if errors.Is(err, usermodels.ErrUserNotFound) {
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

//func Delete(user User) error {
//	err := uservalidation.Validate(uservalidation.User(user), false)
//	if err != nil {
//		return err
//	}
//
//	// ver sobre cache para user continuar logado
//
//}
