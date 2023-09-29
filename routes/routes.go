package routes

import (
	"github.com/jessicamosouza/login-system/handlers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/insert", handlers.CreateUserHandler)
	//http.HandleFunc("/login", controllers.Login)
}
