package sendSms

import (
	"net/http"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

// OK 1- recuperar numero de celular e mensagem da requisição
// OK 2- Validar numero de celular e mensagem
// se qualquer input for invalido retornar http code 400
// 3- fazer chamada no gateway para enviar sms
// se houver erro, retornar erro 500
// se correr tudo bem, retornar http code 201
func ValidateInput(w http.ResponseWriter, r *http.Request) {
	validate = validator.New()

	r.ParseForm()

	phone := r.FormValue("phone")
	message := r.FormValue("mensagem")

	err := validate.Var(phone, "required,number")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Número inválido"))
	}

	err = validate.Var(message, "required,min=1,max=160")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Mensagem dever ter mais de um caracter"))
	}

	send(phone, message, w)
}
