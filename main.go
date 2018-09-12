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
	http.HandleFunc("/home", product.Index)
	http.HandleFunc("/product/show", product.Show)
	http.HandleFunc("/product/addtocart", product.AddToCart)
	http.HandleFunc("/register", user.Register)
	http.HandleFunc("/login", user.Login)
	http.HandleFunc("/logout", user.Logout)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

// SOME USEFUL LINKS
// https://auth0.com/blog/authentication-in-golang/
// https://www.sohamkamani.com/blog/2018/02/25/golang-password-authentication-and-storage/
// https://www.sohamkamani.com/blog/2018/03/25/golang-session-authentication/
// http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go

// For session management
// https://www.objectrocket.com/blog/how-to/top-5-redis-use-cases/
