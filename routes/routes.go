package routes

import (
	"net/http"

	controllers "github.com/jessicamosouza/login-system/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/signup", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)

	// http.HandleFunc("/", controllers.Index)
	// http.HandleFunc("/new", controllers.New)
	// http.HandleFunc("/insert", controllers.Insert)
	// http.HandleFunc("/delete", controllers.Delete)
	// http.HandleFunc("/edit", controllers.Edit)
	// http.HandleFunc("/update", controllers.Update)
}
