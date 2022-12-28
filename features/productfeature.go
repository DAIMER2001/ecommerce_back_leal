package features

import (
	"ecommerce/models"
)

type ProductFeatures struct {
	client models.ProductRepository
}

func NewProductFeatures(repository models.ProductRepository) ProductFeatures {
	return ProductFeatures{client: repository}
}

func (d ProductFeatures) CreateProduct(p models.Product) (err error) {
	return d.client.CreateProduct(p)
}

func (d ProductFeatures) UpdateProduct(p models.Product) (err error) {
	return d.client.UpdateProduct(p)
}

func (d ProductFeatures) DeleteProduct(idProduct string) (err error) {
	return d.client.DeleteProduct(idProduct)
}

func (d ProductFeatures) FindProducts() (m []models.Product, err error) {
	return d.client.FindProducts()
}
