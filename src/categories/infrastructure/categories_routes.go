package infrastructure

import (
	"github.com/gin-gonic/gin"
	"holamundo/src/categories/infrastructure/controllers"
)

func SetupCategoryRoutes(r *gin.Engine, categoryController *controllers.CategoryController, categoryGetController *controllers.CategoryGetController) {
	r.POST("/categories", categoryController.CreateCategory)
	r.GET("/categories", categoryController.ListCategories)
	r.PUT("/categories", categoryController.UpdateCategory)
	r.DELETE("/categories/:id", categoryController.DeleteCategory)
	r.GET("/categories/:id", categoryGetController.GetCategory)
}

