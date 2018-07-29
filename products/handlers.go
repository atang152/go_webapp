package products

import (
	"database/sql"
	"fmt"
	"github.com/go_webapp/config"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	products, err := AllProduct()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "index.html", products)
}

func Show(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	p, err := OneProduct(r)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "product.html", p)

}

func AddToCart(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		//Form submitted
		r.ParseForm()
		fmt.Println(r.Form["product-size"])
		fmt.Println(r.Form["product-color"])
	} else {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
}
