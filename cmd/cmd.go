package cmd

import (
	"fmt"
	"github.com/jessicamosouza/login-system/pkg/routes"
	"log"
	"net/http"
)

func Run() {
	routes.LoadRoutes()
	fmt.Printf("Server is listening on :8000...\n")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
