package main

import (
	"github.com/atang152/go_webapp/product"
	"github.com/atang152/go_webapp/user"
	"net/http"
)

func main() {
	// Add route to serve css files
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.HandleFunc("/", index)
	http.HandleFunc("/product", product.Index)
	http.HandleFunc("/product/show", product.Show)
	http.HandleFunc("/product/addtocart", product.AddToCart)
	http.HandleFunc("/register", user.Register)
	http.HandleFunc("/register/user", user.RegisterUser)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/product", http.StatusSeeOther)
}
