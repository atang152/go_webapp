package user

import (
	"fmt"
	"github.com/atang152/go_webapp/config"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {

	if AlreadyLoggedIn(r) {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	// Insert username
	if r.Method == http.MethodPost {
		u, err := InsertUser(r)

		if err != nil {
			http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
			return
		}

		http.SetCookie(w, u.Cookie)
		fmt.Println(u)

		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
	config.TPL.ExecuteTemplate(w, "register.html", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {

	if AlreadyLoggedIn(r) {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {

		// Get user information and Authenticate
		u, err := GetUser(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		http.SetCookie(w, u.Cookie)

		fmt.Println(u)

		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}

	config.TPL.ExecuteTemplate(w, "login.html", nil)
}

func logout(w http.ResponseWriter, r *http.Request) {

	if !AlreadyLoggedIn(r) {
		http.Redirect(w, req, "/home", http.StatusSeeOther)
	}

	c, err := r.Cookie("session")

	if err != nil {
		http.Error(w, http.StatusText("Internal server error"), http.StatusInternalServerError)
		return
	}

	// Delete the session

}
