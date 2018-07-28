package products

import (
	"errors"
	"github.com/go_webapp/config"
	"net/http"
)

// Specifications and fields for products that we sell on Webapp
type Product struct {
	Id       int
	Type     string
	Name     string
	Url      string
	Imgpath  string
	Price    float32
	Currency string
}

func AllProduct() ([]Product, error) {
	rows, err := config.DB.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := make([]Product, 0)

	for rows.Next() {
		p := Product{}
		err := rows.Scan(&p.Id, &p.Type, &p.Name, &p.Url, &p.Imgpath, &p.Price, &p.Currency)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func OneProduct(r *http.Request) (Product, error) {
	p := Product{}
	name := r.FormValue("name")
	if name == "" {
		return p, errors.New("400. Bad Requests.")
	}

	row := config.DB.QueryRow("SELECT * FROM products where name = $1", name)

	err := row.Scan(&p.Id, &p.Type, &p.Name, &p.Url, &p.Imgpath, &p.Price, &p.Currency)

	if err != nil {
		return p, err
	}

	return p, nil
}
