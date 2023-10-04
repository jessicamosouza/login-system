package managers

import (
	"database/sql"
	"errors"
	"github.com/jessicamosouza/login-system/pkg/db"
	"github.com/jessicamosouza/login-system/pkg/models"
	"github.com/jessicamosouza/login-system/pkg/security"
	"github.com/jessicamosouza/login-system/pkg/validators"
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

	db, err := connectDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	err = models.CreateUser(db, user.FirstName, user.LastName, user.Email, passwordHash)
	if err != nil {
		return err
	}

	return nil
}
func connectDatabase() (*sql.DB, error) {
	db := db.InitDB()
	if db == nil {
		return nil, errors.New("failed to initialize the database connection")
	}
	return db, nil
}
