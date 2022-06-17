package controllers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

	database "github.com/MelkdeSousa/ghopping/database/repository/gorm"
	"github.com/google/uuid"
)

type Product struct {
	Templates  *template.Template
	Repository *database.ProductRepository
}

func loadTemplates() (*template.Template, error) {
	dirs := []string{
		filepath.Join("views", "*.html"),
		filepath.Join("views", "components", "*.html"),
	}

	files := []string{}

	for _, dir := range dirs {
		file, err := filepath.Glob(dir)

		if err != nil {
			return nil, err
		}

		files = append(files, file...)
	}

	t, err := template.ParseFiles(files...)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func NewProductController() *Product {
	templates, err := loadTemplates()

	if err != nil {
		panic(err)
	}

	return &Product{
		Templates:  templates,
		Repository: database.NewProductRepository(),
	}
}

func (p *Product) Index(w http.ResponseWriter, r *http.Request) {
	products, _ := p.Repository.GetAll()
	p.Templates.ExecuteTemplate(w, "index", products)
}

func (p *Product) New(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := uuid.NewString()
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, _ := strconv.Atoi(r.FormValue("price"))
		quantity, _ := strconv.Atoi(r.FormValue("quantity"))

		err := p.Repository.Insert(
			id,
			name,
			description,
			price,
			quantity,
		)

		if err != nil {
			panic(err)
		}
	}

	p.Templates.ExecuteTemplate(w, "new", nil)
}

func (p *Product) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := p.Repository.DeleteById(id)

	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func (p *Product) Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product, err := p.Repository.GetById(id)

	if err != nil {
		panic(err)
	}

	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, _ := strconv.Atoi(r.FormValue("price"))
		quantity, _ := strconv.Atoi(r.FormValue("quantity"))

		err := p.Repository.UpdateById(
			id,
			name,
			description,
			price,
			quantity,
		)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	}

	p.Templates.ExecuteTemplate(w, "edit", product)
}
