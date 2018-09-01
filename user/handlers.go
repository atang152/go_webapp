package user

import (
	"fmt"
	"github.com/atang152/go_webapp/config"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "register.html", nil)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// Insert username
	u, err := InsertUser(r)
	fmt.Println(u)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}
}
