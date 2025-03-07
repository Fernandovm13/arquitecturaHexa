package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"holamundo/src/core" 
	"holamundo/src/products/application"
	"holamundo/src/products/domain/entities"
)

type ProductController struct {
	createUseCase *application.CreateProductUseCase
	listUseCase   *application.ListProductUseCase
	updateUseCase *application.UpdateProductUseCase
	deleteUseCase *application.DeleteProductUseCase
	rabbitMQ      *core.RabbitMQ 
}

func NewProductController(
	create *application.CreateProductUseCase,
	list *application.ListProductUseCase,
	update *application.UpdateProductUseCase,
	deleteUC *application.DeleteProductUseCase,
	rabbitMQ *core.RabbitMQ, 
) *ProductController {
	return &ProductController{
		createUseCase: create,
		listUseCase:   list,
		updateUseCase: update,
		deleteUseCase: deleteUC,
		rabbitMQ:      rabbitMQ, 
	}
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var product entities.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := c.createUseCase.Execute(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create product"})
		return
	}
	ctx.JSON(http.StatusCreated, product)
}

func (c *ProductController) ListProducts(ctx *gin.Context) {
	products, err := c.listUseCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	var product entities.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.updateUseCase.Execute(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (c *ProductController) DeleteProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.deleteUseCase.Execute(int32(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete product"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

func (c *ProductController) BuyProduct(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	message := "Producto con ID " + strconv.Itoa(id) + " ha sido comprado."

	err = c.rabbitMQ.PublishMessage(message)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al enviar mensaje a RabbitMQ"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Compra realizada con éxito y notificación enviada"})
}
