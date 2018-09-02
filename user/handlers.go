package user

import (
	"fmt"
	"github.com/atang152/go_webapp/config"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {

	// if AlreadyLoggedIn(r) {
	// 	http.Redirect(w, r, "/login", http.StatusSeeOther)
	// 	return
	// }

	// Insert username
	if r.Method == http.MethodPost {
		u, err := InsertUser(r)
		http.SetCookie(w, u.Cookie)
		fmt.Println(u)

		if err != nil {
			http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
			return
		}
		http.Redirect(w, r, "/product", http.StatusSeeOther)
	}
	config.TPL.ExecuteTemplate(w, "register.html", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {

	// if AlreadyLoggedIn(r) {
	// 	http.Redirect(w, r, "/home", http.StatusSeeOther)
	// 	return
	// }

	if r.Method == http.MethodPost {
		// Get user information and Authenticate
		u, err := GetUser(r)
		if err != nil {

			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		fmt.Print(u)

		//http.Redirect(w, r, "/product", http.StatusSeeOther)
	}

	config.TPL.ExecuteTemplate(w, "login.html", nil)
}
