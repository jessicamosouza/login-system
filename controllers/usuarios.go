package controllers

import (
	"html/template"
	"log"
	"net/http"

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
		email := r.FormValue("email")
		password := r.FormValue("password")

		models.NewUser(fname, lname, email, password)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
