package main

import (
	"net/http"

	"github.com/MelkdeSousa/ghopping/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8080", nil)
}
