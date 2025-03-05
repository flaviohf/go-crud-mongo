package routes

import (
	"crud-mongo/internal/controllers"
	"crud-mongo/internal/database"
	gateway_impl "crud-mongo/internal/gateways/impl"
	"crud-mongo/internal/gateways/redis"
	"crud-mongo/internal/gateways/repositories"
	"crud-mongo/internal/usecases"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() {
	db := database.GetDataBase()
	redisClient := database.GetRedisConnection()
	redisCache := redis.NewRedisCache(redisClient)

	//repositories
	productRepository := repositories.NewProductRepository(db)

	//gateways
	productGateway := gateway_impl.NewProductGateway(productRepository, redisCache)

	//usecases
	productUsecase := usecases.NewProductUsecase(productGateway)

	//controllers
	productController := controllers.NewProductController(productUsecase)

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/mongo/api/v1/products", productController.CreateProduct)
	router.GET("/mongo/api/v1/products", productController.GetProducts)
	router.GET("/mongo/api/v1/products/:id", productController.GetProductByID)
	router.PUT("/mongo/api/v1/products/:id", productController.UpdateProduct)
	router.DELETE("/mongo/api/v1/products/:id", productController.DeleteProduct)
	router.Run(":8080")
}
