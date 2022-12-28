package models

import "github.com/google/uuid"

type Product struct {
	IDProduct          uuid.UUID `json:"id_product,omitempty"`
	Name               string    `json:"name"`
	Price              int64     `json:"price"`
	AccumulationPoints int64     `json:"acumulation_points"`
	Image              string    `json:"image"`
	Description        string    `json:"description"`
}

type ProductRepository interface {
	CreateProduct(product Product) (err error)
	UpdateProduct(product Product) (err error)
	FindProducts() (m []Product, err error)
	DeleteProduct(idProduct string) (err error)
}
