package managers

import (
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
	err := validators.Validate(validators.User(User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}))
	if err != nil {
		return err
	}

	_, err = security.GenerateHash(user.Password)
	if err != nil {
		return err
	}

	// Perform user creation in the database (you may need to implement this)
	// Example:
	// err = db.CreateUser(user.FirstName, user.LastName, user.Email, passwordHash)
	// if err != nil {
	//     return err
	// }
	return nil
}
