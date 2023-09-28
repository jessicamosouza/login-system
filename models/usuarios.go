package models

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jessicamosouza/login-system/db"
)

var ErrUserNotFound = errors.New("[models] user not found")

type User struct {
	FirstName string `json:"fname" db:"firstname"`
	Lastname  string `json:"lname" db:"lastname"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
}

func InsertUser(firstName, lastName, email, password string) error {
	db := db.InitDB()
	defer db.Close()

	addUserDB, err := db.Prepare("insert into users(firstname, lastname, email, password)  values($1,$2,$3,$4)")
	if err != nil {
		return fmt.Errorf("[models] error preparing insert: %w", err)
	}

	addUserDB.Exec(firstName, lastName, email, password)
	if err != nil {
		return fmt.Errorf("[models] error insert: %w", err)
	}

	return nil
}

func GetUser(email, password string) (string, error) {
	db := db.InitDB()
	defer db.Close()

	getUserDB := db.QueryRow("SELECT password FROM users WHERE email LIKE $1;", email)
	storedUserData := &User{}

	err := getUserDB.Scan(&storedUserData.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrUserNotFound
		}
		return "", fmt.Errorf("[models] error finding user password: %w", err)
	}
	println("stored: ", storedUserData.Password)

	return storedUserData.Password, nil
}
