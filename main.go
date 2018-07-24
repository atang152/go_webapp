package main

import (
	// "encoding/json"
	// "fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

// type productSpecs []struct {
// 	Id           string `json:"id"`
// 	URL          string `json:"url"`
// 	Img          string `json:"img"`
// 	ProductName  string `json:"product_name"`
// 	ProductPrice string `json:"product_price"`
// }

func main() {

	// var data []productSpecs

	// rcvd := `[{
	//     "id": "1",
	//     "url": "fancy-cardigan.html",
	//     "img": "../static/img/photo_1.jpg",
	//     "product_name": "Fancy Cardigan",
	//     "product_price": "HKD 399"
	// }, {
	//     "id": "2",
	//     "url": "fancy-cardigan.html",
	//     "img": "../static/img/photo_2.jpg",
	//     "product_name": "Fancy Cardigan",
	//     "product_price": "HKD 399"
	// }]`

	// err := json.Unmarshal([]byte(rcvd), &data)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// for _, v := range data {
	// 	fmt.Println(v.ProductName)
	// }

	// Add route to serve css files
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.HandleFunc("/", index)
	http.HandleFunc("/fancy-cardigan.html", product)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatalln("error retrieving template: ", err)
	}
}

func product(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "fancy-cardigan.html", nil)
	if err != nil {
		log.Fatalln("error retrieving template: ", err)
	}
}
