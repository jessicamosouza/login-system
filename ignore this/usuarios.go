package ignore_this

import (
	"errors"
	"fmt"
	"github.com/jessicamosouza/login-system/pkg/models"
	"io"
	"net/http"
	"net/mail"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	FirstName string `json:"firstName" validate:"required,alpha,min=2,max=100"`
	LastName  string `json:"lastName" validate:"required,alpha,min=2,max=100"`
	Email     string `json:"email" validate:"required,unique=email,email"`
	Password  string `json:"password" validate:"required,min=8,password"`
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

func GetBody(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userData, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	close(r)

	fmt.Println(string(userData))

	_, err = w.Write([]byte("Ok!"))
	if err != nil {
		return
	}

	//models.InsertUser(u.FirstName, u.LastName, u.Email, passwordHash)
	//
	//// mensagem de registrado com sucesso, então redirecionar para login ou pagina inicial
	//http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func close(r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(r.Body)
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
