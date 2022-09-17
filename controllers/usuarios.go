package controllers

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"net/mail"

	"github.com/go-passwd/validator"
	"github.com/jessicamosouza/login-system/models"
	"golang.org/x/crypto/bcrypt"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allUsers := models.SearchAllUsers()
	err := temp.ExecuteTemplate(w, "Index", allUsers)
	if err != nil {
		log.Println(err)
	}
}

func New(w http.ResponseWriter, r *http.Request) {
	err := temp.ExecuteTemplate(w, "New", nil)
	if err != nil {
		log.Print(err)
	}
}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		fname := r.FormValue("fname")
		err := checkName(fname)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("First name must contain at least 2 characters."))
			return
		}

		lname := r.FormValue("lname")
		err = checkName(lname)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Last name must contain at least 2 characters."))
			return
		}

		email := r.FormValue("email")
		err = checkEmail(email)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid email."))
			return
		}

		password := r.FormValue("password")
		err = checkPassword(password)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		passwordHash, err := generateHash(password)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		models.InsertUser(fname, lname, email, passwordHash)
	}

	// mensagem de registrado com sucesso, então redirecionar para login ou pagina inicial
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func checkName(name string) error {

	if len(name) > 2 {
		return nil
	}
	return errors.New("invalid name")

}

func checkEmail(email string) error {
	_, err := mail.ParseAddress(email)

	return err
}

func checkPassword(password string) error {
	passwordValidator := validator.New(validator.Regex("^&&w+$", errors.New("invalid password")))

	return passwordValidator.Validate(password)

}

func generateHash(password string) (string, error) {
	const cost = 14
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	return string(passwordHash), err
}
