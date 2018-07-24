package models

import "math/big"

// Specifications and fields for products that we sell on Webapp
type Product struct {
	Url   string
	Image string
	Name  string
	Price *big.Float
}

// Returns a new Product instances
func NewProduct(url, image string, name string, price *big.Float) *Product {
	return &Product{
		Url:   url,
		Image: image,
		Name:  name,
		Price: price,
	}
}
