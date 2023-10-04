package managers

import (
	"database/sql"
	"errors"
	"github.com/jessicamosouza/login-system/pkg/db"
	"github.com/jessicamosouza/login-system/pkg/models"
	"github.com/jessicamosouza/login-system/pkg/security"
	"github.com/jessicamosouza/login-system/pkg/validators"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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

// Login TODO: refactor this function
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

	l := User{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	hashPassword, err := models.GetUser(l.Email, l.Password)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if checkPasswordHash(l.Password, hashPassword) {
		http.Redirect(w, r, "/welcome", http.StatusMovedPermanently)
	}

}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func connectDatabase() (*sql.DB, error) {
	db := db.InitDB()
	if db == nil {
		return nil, errors.New("failed to initialize the database connection")
	}
	return db, nil
}
