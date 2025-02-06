package main

import (
	"github.com/gin-contrib/cors" // <-- Importa el middleware de CORS
	"github.com/gin-gonic/gin"
	"holamundo/src/core"

	productApp "holamundo/src/products/application"
	productRepo "holamundo/src/products/infrastructure/repositories"
	productCtrl "holamundo/src/products/infrastructure/controllers"
	productGetCtrl "holamundo/src/products/infrastructure/controllers"

	categoryApp "holamundo/src/categories/application"
	categoryRepo "holamundo/src/categories/infrastructure/repositories"
	categoryCtrl "holamundo/src/categories/infrastructure/controllers"
	categoryGetCtrl "holamundo/src/categories/infrastructure/controllers"

	productInfra "holamundo/src/products/infrastructure"
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

	getProductUC := productApp.NewGetProductUseCase(pRepo)
	productGetController := productGetCtrl.NewProductGetController(getProductUC)

	cRepo := categoryRepo.NewMySQLCategoryRepository()
	createCategoryUC := categoryApp.NewCreateCategoryUseCase(cRepo)
	listCategoryUC := categoryApp.NewListCategoryUseCase(cRepo)
	updateCategoryUC := categoryApp.NewUpdateCategoryUseCase(cRepo)
	deleteCategoryUC := categoryApp.NewDeleteCategoryUseCase(cRepo)
	cController := categoryCtrl.NewCategoryController(createCategoryUC, listCategoryUC, updateCategoryUC, deleteCategoryUC)

	getCategoryUC := categoryApp.NewGetCategoryUseCase(cRepo)
	categoryGetController := categoryGetCtrl.NewCategoryGetController(getCategoryUC)

	r := gin.Default()

	// Habilitar CORS con la configuración por defecto
	r.Use(cors.Default())

	productInfra.SetupProductRoutes(r, pController, productGetController)
	categoryInfra.SetupCategoryRoutes(r, cController, categoryGetController)

	r.Run(":8080")
}
