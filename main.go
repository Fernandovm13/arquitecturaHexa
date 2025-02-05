package main

import (
	"github.com/gin-gonic/gin"
	
	"holamundo/src/core"
	
	productApp "holamundo/src/products/application"
	productRepo "holamundo/src/products/infrastructure/repositories"
	productCtrl "holamundo/src/products/infrastructure/controllers"
	productInfra "holamundo/src/products/infrastructure" 

	categoryApp "holamundo/src/categories/application"
	categoryRepo "holamundo/src/categories/infrastructure/repositories"
	categoryCtrl "holamundo/src/categories/infrastructure/controllers"
	categoryInfra "holamundo/src/categories/infrastructure"  
)

func main() {
	core.InitDB()
	defer core.CloseDB()

	pRepo := productRepo.NewMySQLProductRepository()
	createProductUC := productApp.NewCreateProductUseCase(pRepo)
	listProductUC := productApp.NewListProductUseCase(pRepo)
	updateProductUC := productApp.NewUpdateProductUseCase(pRepo)
	deleteProductUC := productApp.NewDeleteProductUseCase(pRepo)
	pController := productCtrl.NewProductController(createProductUC, listProductUC, updateProductUC, deleteProductUC)

	cRepo := categoryRepo.NewMySQLCategoryRepository()
	createCategoryUC := categoryApp.NewCreateCategoryUseCase(cRepo)
	listCategoryUC := categoryApp.NewListCategoryUseCase(cRepo)
	updateCategoryUC := categoryApp.NewUpdateCategoryUseCase(cRepo)
	deleteCategoryUC := categoryApp.NewDeleteCategoryUseCase(cRepo)
	cController := categoryCtrl.NewCategoryController(createCategoryUC, listCategoryUC, updateCategoryUC, deleteCategoryUC)

	r := gin.Default()

	productInfra.SetupProductRoutes(r, pController)
	categoryInfra.SetupCategoryRoutes(r, cController)

	r.Run(":8080")
}
