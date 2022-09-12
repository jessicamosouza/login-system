package main

import (
	"log"
	"net/http"

	"github.com/jessicamosouza/login-system/routes"
)

func main() {
	routes.CarregaRotas()
	log.Fatal(http.ListenAndServe(":8000", nil))

}
