package routes

import (
	handlers2 "github.com/jessicamosouza/login-system/http"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/insert", handlers2.CreateUserHandler)
	http.HandleFunc("/login", handlers2.LoginUserHandler)
	http.HandleFunc("/delete", handlers2.DeleteUserHandler)
}
