package main

import (
	"html/template"
	"net/http"
	"path"

	"github.com/MelkdeSousa/ghopping/domain"
)

var TEMPLATES = template.Must(template.ParseGlob(path.Join("views", "*.html")))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []domain.Product{
		{ID: 1, Name: "Laptop", Description: "Laptop description", Price: 1000, Quantity: 134},
		{ID: 2, Name: "Mouse", Description: "Mouse description", Price: 100, Quantity: 167},
		{ID: 3, Name: "Keyboard", Description: "Keyboard description", Price: 200, Quantity: 178},
	}

	TEMPLATES.ExecuteTemplate(w, "index", products)
}
