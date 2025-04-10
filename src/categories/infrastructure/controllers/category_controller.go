package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"holamundo/src/categories/application"
	"holamundo/src/categories/domain/entities"
)

type CategoryController struct {
	createUC *application.CreateCategoryUseCase
	listUC   *application.ListCategoryUseCase
	updateUC *application.UpdateCategoryUseCase
	deleteUC *application.DeleteCategoryUseCase
}

func NewCategoryController(
	createUC *application.CreateCategoryUseCase,
	listUC *application.ListCategoryUseCase,
	updateUC *application.UpdateCategoryUseCase,
	deleteUC *application.DeleteCategoryUseCase,
) *CategoryController {
	return &CategoryController{
		createUC: createUC,
		listUC:   listUC,
		updateUC: updateUC,
		deleteUC: deleteUC,
	}
}

func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	var input struct {
		Name   string `json:"name"`
		Secret string `json:"secret"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if input.Name == "" || input.Secret == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nombre y secreto son obligatorios"})
		return
	}

	category := entities.Category{
		Name: input.Name,
		Secret: sql.NullString{
			String: input.Secret,
			Valid:  true,
		},
	}

	err := c.createUC.Execute(&category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Categoría creada correctamente"})
}

func (c *CategoryController) ListCategories(ctx *gin.Context) {
	categories, err := c.listUC.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var result []entities.CategoryDTO
	for _, cat := range categories {
		result = append(result, entities.CategoryDTO{
			ID:   cat.ID,
			Name: cat.Name,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *CategoryController) UpdateCategory(ctx *gin.Context) {
	var input struct {
		ID     int32  `json:"id"`
		Name   string `json:"name"`
		Secret string `json:"secret"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	category := entities.Category{
		ID:   input.ID,
		Name: input.Name,
		Secret: sql.NullString{
			String: input.Secret,
			Valid:  input.Secret != "",
		},
	}

	err := c.updateUC.Execute(&category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Categoría actualizada correctamente"})
}

func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = c.deleteUC.Execute(int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Categoría eliminada correctamente"})
}
