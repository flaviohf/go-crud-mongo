package gateway_impl

import (
	"crud-mongo/internal/domains"
	"crud-mongo/internal/gateways/redis"
	"crud-mongo/internal/gateways/repositories"
	"encoding/json"
	"log/slog"
	"time"
)

type productGatewayImpl struct {
	repository *repositories.ProductRepository
	redisCache *redis.RedisCache
}

func NewProductGateway(repository *repositories.ProductRepository, redisCache *redis.RedisCache) *productGatewayImpl {
	return &productGatewayImpl{repository: repository, redisCache: redisCache}
}

func (gateway *productGatewayImpl) CreateProduct(product domains.Product) (domains.Product, error) {
	savedProduct, err := gateway.repository.CreateProduct(product)

	if err != nil {
		return domains.Product{}, err
	}

	return savedProduct, nil
}

func (gateway *productGatewayImpl) GetProducts() ([]domains.Product, error) {
	cacheResult, err := gateway.redisCache.GetCache("products")
	if err == nil {
		var products = []domains.Product{}
		json.Unmarshal([]byte(cacheResult), &products)
		slog.Info("Retornando do cache")
		return products, nil
	}

	products, err := gateway.repository.GetProducts()
	if err != nil {
		slog.Error("Erro ao buscar no mongodb")
		return []domains.Product{}, err
	}

	json, err := json.Marshal(products)
	if err != nil {
		slog.Error("Erro para converter para json para salvar no cache")
		return []domains.Product{}, err
	}

	errRedis := gateway.redisCache.SetCache("products", string(json), 180*time.Second)
	if errRedis != nil {
		slog.Error("Erro ao tentar salvar no redis cache")
		return []domains.Product{}, err
	}

	slog.Info("Retornando a quente")
	return products, nil
}

func (gateway *productGatewayImpl) GetProductByID(id string) (domains.Product, error) {
	product, err := gateway.repository.GetProductByID(id)

	if err != nil {
		return domains.Product{}, err
	}

	return product, nil
}

func (gateway *productGatewayImpl) UpdateProduct(product domains.Product) (domains.Product, error) {
	updatedProduct, err := gateway.repository.UpdateProduct(product)

	if err != nil {
		return domains.Product{}, err
	}

	return updatedProduct, nil
}

func (gateway *productGatewayImpl) DeleteProduct(id string) error {
	err := gateway.repository.DeleteProduct(id)

	if err != nil {
		return err
	}

	return nil
}
