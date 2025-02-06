package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "holamundo/src/categories/application"
)

type CategoryGetController struct {
    getUseCase *application.GetCategoryUseCase
}

func NewCategoryGetController(getUseCase *application.GetCategoryUseCase) *CategoryGetController {
    return &CategoryGetController{getUseCase: getUseCase}
}

func (c *CategoryGetController) GetCategory(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de categoría inválido"})
        return
    }

    category, err := c.getUseCase.Execute(int32(id))
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Categoría no encontrada"})
        return
    }

    ctx.JSON(http.StatusOK, category)
}
