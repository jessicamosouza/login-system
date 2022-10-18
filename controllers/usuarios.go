package controllers

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/mail"
	"unicode"

	"github.com/jessicamosouza/login-system/models"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	FirstName string `json:"fname" validate:"required,alpha,min=2,max=100"`
	LastName  string `json:"lname" validate:"required,alpha,min=2,max=100"`
	Email     string `json:"email" validate:"required,unique=email,email"`
	Password  string `json:"password" validate:"required,min=8,password"`
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	err := temp.ExecuteTemplate(w, "Index", nil)
	if err != nil {
		log.Println(err)
	}
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	err := temp.ExecuteTemplate(w, "Welcome", nil)
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

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// mensagem de registrado com sucesso, então redirecionar para login ou pagina inicial
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

	l := User{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	result, err := models.GetUser(l.Email, l.Password)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println(CheckPasswordHash(l.Password, result), result, err)
	
	if CheckPasswordHash(l.Password, result) {
		http.Redirect(w, r, "/welcome", http.StatusMovedPermanently)
	}

}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// mensagem de registrado com sucesso, então redirecionar para login ou pagina inicial
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

	u := User{
		FirstName: r.FormValue("fname"),
		LastName:  r.FormValue("lname"),
		Email:     r.FormValue("email"),
		Password:  r.FormValue("password"),
	}

	if err := checkUser(&u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	passwordHash, err := generateHash(u.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	models.InsertUser(u.FirstName, u.LastName, u.Email, passwordHash)

	// mensagem de registrado com sucesso, então redirecionar para login ou pagina inicial
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

// Handlers (Controller) -> Negocios (UseCases) -> Repositório (Acesso a dados (models))

func checkUser(u *User) error {
	if checkName(u.FirstName) != nil {
		return errors.New("first name must contain at least 2 characters")
	}
	if checkName(u.LastName) != nil {
		return errors.New("last name must contain at least 2 characters")
	}

	if checkEmail(u.Email) != nil {
		return errors.New("invalid email")
	}

	if checkPassword(u.Password) != nil {
		return errors.New("invalid password")
	}
	return nil
}

func checkName(name string) error {
	if len(name) <= 2 {
		return errors.New("invalid name")
	}
	return nil
}

func checkEmail(email string) error {
	mail, err := mail.ParseAddress(email)
	fmt.Println(mail, err)
	return err
}

func checkPassword(password string) error {
	// upp: at least one upper case letter.
	// low: at least one lower case letter.
	// num: at least one digit.
	// sym: at least one special character.
	// tot: at least eight characters long.
	// No empty string or whitespace.

	var (
		upp, low, num, sym bool
		tot                uint8
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return errors.New("invalid password")
		}
	}

	if !upp || !low || !num || !sym || tot < 8 {
		return errors.New("invalid password")
	}

	return nil
}

func generateHash(password string) (string, error) {
	const cost = 14
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	return string(passwordHash), err
}
