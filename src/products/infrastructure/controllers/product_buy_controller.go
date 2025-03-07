package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "holamundo/src/core" 
    "holamundo/src/products/application"
)

type ProductBuyController struct {
    buyUseCase *application.BuyProductUseCase
}

func NewProductBuyController(buyUseCase *application.BuyProductUseCase) *ProductBuyController {
    return &ProductBuyController{buyUseCase: buyUseCase}
}

func (c *ProductBuyController) BuyProduct(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de producto inválido"})
        return
    }

    rabbitMQ, err := core.NewRabbitMQ("hello") 
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al conectar con RabbitMQ"})
        return
    }
    defer rabbitMQ.Close() 

    if err := c.buyUseCase.Execute(int32(id)); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Compra realizada y notificación enviada"})
}