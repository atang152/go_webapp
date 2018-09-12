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

func Logout(w http.ResponseWriter, r *http.Request) {

	// If user is not logged in; redirect back to home
	if !AlreadyLoggedIn(r) {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}

	// Delete Session in Database
	err := DeleteSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Delete Cookie
	c := &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}

	http.SetCookie(w, c)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
