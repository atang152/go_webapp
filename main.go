package main

import (
	"html/template"
	"log"
	"math/big"
	"net/http"

	"github.com/go_webapp/models"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	// Add route to serve css files
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.HandleFunc("/", index)
	http.HandleFunc("/products.html", product)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	// Test Database
	data := map[string]interface{}{
		"Products": []*models.Product{
			models.NewProduct("products.html", "../static/img/photo_1.jpg", "Fancy Cardigan", big.NewFloat(399)),
			models.NewProduct("products.html", "../static/img/photo_2.jpg", "Fancy Dress", big.NewFloat(599)),
			models.NewProduct("products.html", "../static/img/photo_3.jpg", "Fancy Coat", big.NewFloat(1888)),
			models.NewProduct("products.html", "../static/img/photo_4.jpg", "Fancy Shirt", big.NewFloat(299)),
		},
	}

	// Execute template and pass data in
	err := tpl.ExecuteTemplate(w, "index.html", data)

	if err != nil {
		log.Fatalln("error retrieving template: ", err)
	}
}

func product(w http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(w, "products.html", nil)
	if err != nil {
		log.Fatalln("error retrieving template: ", err)
	}
}
