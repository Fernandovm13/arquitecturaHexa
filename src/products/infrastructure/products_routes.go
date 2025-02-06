package infrastructure

import (
	"github.com/gin-gonic/gin"
	"holamundo/src/products/infrastructure/controllers"
)

func SetupProductRoutes(r *gin.Engine, productController *controllers.ProductController, productGetController *controllers.ProductGetController) {
	r.POST("/products", productController.CreateProduct)
	r.GET("/products", productController.ListProducts)
	r.PUT("/products", productController.UpdateProduct)
	r.DELETE("/products/:id", productController.DeleteProduct)
	r.GET("/products/:id", productGetController.GetProduct)
}
