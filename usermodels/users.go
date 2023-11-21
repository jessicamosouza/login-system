package usermodels

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jessicamosouza/login-system/db"
)

var ErrUserNotFound = errors.New("[usermodels] user not found")

type User struct {
	FirstName string `json:"fname" db:"firstname"`
	Lastname  string `json:"lname" db:"lastname"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
}

func CreateUser(firstName, lastName, email, password string) error {
	db := db.InitDB()
	defer db.Close()

	addUserDB, err := db.Prepare("INSERT INTO \"User\" (first_name, last_name, email, password)  values($1,$2,$3,$4)")
	if err != nil {
		return fmt.Errorf("[usermodels] error preparing insert: %w", err)
	}

	addUserDB.Exec(firstName, lastName, email, password)
	if err != nil {
		return fmt.Errorf("[usermodels] error insert: %w", err)
	}

	return nil
}

func GetUser(email, password string) (string, error) {
	db := db.InitDB()
	defer db.Close()

	getUserDB := db.QueryRow("SELECT password FROM \"User\" WHERE email LIKE $1;", email)
	storedUserData := &User{}

	err := getUserDB.Scan(&storedUserData.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrUserNotFound
		}
		return "", fmt.Errorf("[usermodels] error finding user password: %w", err)
	}
	println("stored: ", storedUserData.Password)

	return storedUserData.Password, nil
}
