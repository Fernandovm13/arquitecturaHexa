package infrastructure

import (
	"github.com/gin-gonic/gin"
	"holamundo/src/categories/infrastructure/controllers"
)

func SetupCategoryRoutes(r *gin.Engine, cController *controllers.CategoryController, categoryGetController *controllers.CategoryGetController) {
	r.POST("/categories", cController.CreateCategory)       
	r.GET("/categories", cController.ListCategories)         
	r.GET("/categories/:id", categoryGetController.GetCategory) 
	r.PUT("/categories", cController.UpdateCategory)        
	r.DELETE("/categories/:id", cController.DeleteCategory) 
}
