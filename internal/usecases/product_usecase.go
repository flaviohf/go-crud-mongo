package usecases

import (
	"crud-mongo/internal/domains"
	"crud-mongo/internal/gateways"
	"errors"
	"log/slog"
)

type ProductUsecase struct {
	gateway gateways.ProductGateway
}

func NewProductUsecase(gateway gateways.ProductGateway) *ProductUsecase {
	return &ProductUsecase{gateway: gateway}
}

func (usecase *ProductUsecase) CreateProduct(product domains.Product) (domains.Product, error) {
	createdProduct, err := usecase.gateway.CreateProduct(product)

	if err != nil {
		slog.Error("Error to create product", "erro", err)
		return domains.Product{}, err
	}

	return createdProduct, nil
}

func (usecase *ProductUsecase) GetProducts() ([]domains.Product, error) {
	products, err := usecase.gateway.GetProducts()

	if err != nil {
		slog.Error("Error to get products", "erro", err)
		return nil, err
	}

	return products, nil
}

func (usecase *ProductUsecase) GetProductByID(id string) (domains.Product, error) {
	foundProduct, err := usecase.gateway.GetProductByID(id)

	if err != nil {
		slog.Error("Error to get product", "erro", err)
		return domains.Product{}, err
	}

	if foundProduct.ID == "" {
		slog.Error("Product not found", "erro", err)
		return domains.Product{}, errors.New("product not found")
	}

	return foundProduct, nil
}

func (usecase *ProductUsecase) UpdateProduct(product domains.Product, id string) (domains.Product, error) {
	_, err := usecase.GetProductByID(id)

	if err != nil && err.Error() == "product not found" {
		return domains.Product{}, errors.New("product not found")
	}

	product.ID = id
	updatedProduct, err := usecase.gateway.UpdateProduct(product)

	if err != nil {
		slog.Error("Error to update product", "erro", err)
		return domains.Product{}, err
	}

	return updatedProduct, nil
}

func (usecase *ProductUsecase) DeleteProduct(id string) error {
	err := usecase.gateway.DeleteProduct(id)

	if err != nil {
		slog.Error("Error to delete product", "erro", err)
		return err
	}

	return nil
}
