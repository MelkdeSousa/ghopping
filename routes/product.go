package routes

import (
	"net/http"

	"github.com/MelkdeSousa/ghopping/controllers"
)

func LoadRoutes() {
	productController := controllers.NewProductController()

	http.HandleFunc("/", productController.Index)
	http.HandleFunc("/new", productController.New)
	http.HandleFunc("/delete", productController.Delete)
	http.HandleFunc("/edit", productController.Edit)
}
