package main

import (
	"html/template"
	"net/http"
	"path"
)

var TEMPLATES = template.Must(template.ParseGlob(path.Join("views", "*.html")))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	TEMPLATES.ExecuteTemplate(w, "index", nil)
}
