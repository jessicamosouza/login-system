package routes

import (
	"github.com/jessicamosouza/login-system/pkg/handlers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/insert", handlers.CreateUserHandler)
	//http.HandleFunc("/login", ignore this.Login)
}
