package gateways

import "crud-mongo/internal/domains"

type ProductGateway interface {
	CreateProduct(product domains.Product) (domains.Product, error)
	GetProducts() ([]domains.Product, error)
	GetProductByID(id string) (domains.Product, error)
	UpdateProduct(product domains.Product) (domains.Product, error)
	DeleteProduct(id string) error
}
