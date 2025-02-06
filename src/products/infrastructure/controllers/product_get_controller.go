package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "holamundo/src/products/application"
)

type ProductGetController struct {
    getUseCase *application.GetProductUseCase
}

func NewProductGetController(getUseCase *application.GetProductUseCase) *ProductGetController {
    return &ProductGetController{getUseCase: getUseCase}
}

func (c *ProductGetController) GetProduct(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de producto inválido"})
        return
    }

    product, err := c.getUseCase.Execute(int32(id))
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
        return
    }

    ctx.JSON(http.StatusOK, product)
}
