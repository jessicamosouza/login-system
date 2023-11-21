package routes

import (
	handler "github.com/jessicamosouza/login-system/http"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/insert", handler.CreateUserHandler)
	http.HandleFunc("/login", handler.LoginUserHandler)
	http.HandleFunc("/delete", handler.DeleteUserHandler)
}
