package routes

import (
	"net/http"

	controllers "github.com/jessicamosouza/login-system/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/signup", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/welcome", controllers.Welcome)
}
