package repositories

import (
	"context"
	"crud-mongo/internal/domains"
	"log/slog"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	db *mongo.Database
}

func NewProductRepository(db *mongo.Database) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repository *ProductRepository) CreateProduct(product domains.Product) (domains.Product, error) {
	product.ID = uuid.New().String()
	result, err := repository.db.Collection("products").InsertOne(context.Background(), product)

	if err != nil {
		slog.Error("Error to get products", "erro", err)
		return domains.Product{}, err
	}

	if id, ok := result.InsertedID.(string); ok {
		product.ID = id
	} else {
		slog.Error("Error to assert InsertedID to string", "erro", err)
		return domains.Product{}, err
	}
	return product, nil
}

func (repository *ProductRepository) GetProducts() ([]domains.Product, error) {
	products := []domains.Product{}
	cursor, err := repository.db.Collection("products").Find(context.Background(), bson.M{})

	if err != nil {
		slog.Error("Error to get products", "erro", err)
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		product := domains.Product{}
		if err := cursor.Decode(&product); err != nil {
			slog.Error("Error to decode product", "erro", err)
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (repository *ProductRepository) GetProductByID(id string) (domains.Product, error) {
	product := domains.Product{}
	err := repository.db.Collection("products").FindOne(context.Background(), bson.M{"_id": id}).Decode(&product)

	if err != nil {
		slog.Error("Error to get product by id", "erro", err)
		return domains.Product{}, err
	}

	return product, nil
}

func (repository *ProductRepository) UpdateProduct(product domains.Product) (domains.Product, error) {
	_, err := repository.db.Collection("products").UpdateOne(context.Background(), bson.M{"_id": product.ID}, bson.M{"$set": product})

	if err != nil {
		slog.Error("Error to update product", "erro", err)
		return domains.Product{}, err
	}

	return product, nil
}

func (repository *ProductRepository) DeleteProduct(id string) error {
	_, err := repository.db.Collection("products").DeleteOne(context.Background(), bson.M{"_id": id})

	if err != nil {
		slog.Error("Error to delete product", "erro", err)
		return err
	}

	return nil
}
