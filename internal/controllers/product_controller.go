package controllers

import (
	"crud-mongo/internal/domains"
	"crud-mongo/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	usecase *usecases.ProductUsecase
}

func NewProductController(usecase *usecases.ProductUsecase) *productController {
	return &productController{usecase: usecase}
}

// CreateProduct cria um novo produto
// @Summary cria um produto
// @Description Cria um novo produto
// @Tags Products
// @Accept json
// @Produce json
// @Param product body domains.Product true "Dados do produto"
// @Success 201 {object} domains.Product
// @Router /mongo/api/v1/products/:id [post]
func (controller *productController) CreateProduct(context *gin.Context) {
	var product domains.Product

	if err := context.ShouldBindJSON(&product); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProduct, err := controller.usecase.CreateProduct(product)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, createdProduct)
}

// GetProducts retorna uma lista de produtos
// @Summary Lista produtos
// @Description Retorna todos os produtos cadastrados
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {array} domains.Product
// @Router /mongo/api/v1/products [get]
func (controller *productController) GetProducts(context *gin.Context) {
	products, err := controller.usecase.GetProducts()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, products)
}

// GetProductByID retorna um produto pelo seu ID
// @Summary retorna um produto
// @Description Retorna um produto específico pelo seu ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "ID do produto"
// @Success 200 {object} domains.Product
// @Router /mongo/api/v1/products/:id [get]
func (controller *productController) GetProductByID(context *gin.Context) {
	productId := context.Param("id")

	product, err := controller.usecase.GetProductByID(productId)

	if err != nil {
		if err.Error() == "product not found" {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	context.JSON(http.StatusOK, product)
}

// UpdateProduct atualiza um produto pelo seu ID
// @Summary atualiza um produto
// @Description Atualiza um produto específico pelo seu ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "ID do produto"
// @Param product body domains.Product true "Dados do produto"
// @Success 200 {object} domains.Product
// @Router /mongo/api/v1/products/:id [put]
func (controller *productController) UpdateProduct(context *gin.Context) {
	productId := context.Param("id")
	var product domains.Product

	if err := context.ShouldBindJSON(&product); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProduct, err := controller.usecase.UpdateProduct(product, productId)

	if err != nil {
		if err.Error() == "product not found" {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	context.JSON(http.StatusOK, updatedProduct)
}

// DeleteProduct deleta um produto pelo seu ID
// @Summary deleta um produto
// @Description Deleta um produto específico pelo seu ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "ID do produto"
// @Success 204
// @Router /mongo/api/v1/products/:id [delete]
func (controller *productController) DeleteProduct(context *gin.Context) {
	productId := context.Param("id")

	err := controller.usecase.DeleteProduct(productId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, nil)
}
