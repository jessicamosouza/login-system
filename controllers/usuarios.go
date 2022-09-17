package controllers

import (
	"html/template"
	"log"
	"net/http"
	"net/mail"

	"github.com/jessicamosouza/login-system/models"
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
		lname := r.FormValue("lname")

		email := checkEmail(w, r)

		password := r.FormValue("password")

		models.InsertUser(fname, lname, password, email)
	}

	// mensagem de registrado com sucesso, então redirecionar para login ou pagina inicial
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func checkEmail(w http.ResponseWriter, r *http.Request) *mail.Address {
	email, err := mail.ParseAddress(r.FormValue("email"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid email"))
	}

	return email
}
