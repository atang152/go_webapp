package main

import (
	"github.com/go_webapp/products"
	"net/http"
)

func main() {
	// Add route to serve css files
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.HandleFunc("/", index)
	http.HandleFunc("/products", products.Index)
	http.HandleFunc("/products/show", products.Show)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/products", http.StatusSeeOther)
}
