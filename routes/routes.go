package routes

import (
	"github.com/jessicamosouza/login-system/handlers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/insert", handlers.GetUserData)
	//http.HandleFunc("/login", controllers.Login)
}
