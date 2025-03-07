package main

import (
	"log"
	"github.com/gin-contrib/cors" 
	"github.com/gin-gonic/gin"
	"holamundo/src/core"

	productApp "holamundo/src/products/application"
	productRepo "holamundo/src/products/infrastructure/repositories"
	productCtrl "holamundo/src/products/infrastructure/controllers"
	productGetCtrl "holamundo/src/products/infrastructure/controllers"
	productBuyCtrl "holamundo/src/products/infrastructure/controllers"

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

	rabbitMQ, err := core.NewRabbitMQ("compras")
	if err != nil {
		log.Fatalf("Error al conectar con RabbitMQ: %v", err)
	}
	defer rabbitMQ.Close()

	pRepo := productRepo.NewMySQLProductRepository()
	createProductUC := productApp.NewCreateProductUseCase(pRepo)
	listProductUC := productApp.NewListProductUseCase(pRepo)
	updateProductUC := productApp.NewUpdateProductUseCase(pRepo)
	deleteProductUC := productApp.NewDeleteProductUseCase(pRepo)

	pController := productCtrl.NewProductController(createProductUC, listProductUC, updateProductUC, deleteProductUC, rabbitMQ)

	getProductUC := productApp.NewGetProductUseCase(pRepo)
	productGetController := productGetCtrl.NewProductGetController(getProductUC)

	buyProductUC := productApp.NewBuyProductUseCase(pRepo, rabbitMQ)
	productBuyController := productBuyCtrl.NewProductBuyController(buyProductUC)

	cRepo := categoryRepo.NewMySQLCategoryRepository()
	createCategoryUC := categoryApp.NewCreateCategoryUseCase(cRepo)
	listCategoryUC := categoryApp.NewListCategoryUseCase(cRepo)
	updateCategoryUC := categoryApp.NewUpdateCategoryUseCase(cRepo)
	deleteCategoryUC := categoryApp.NewDeleteCategoryUseCase(cRepo)
	cController := categoryCtrl.NewCategoryController(createCategoryUC, listCategoryUC, updateCategoryUC, deleteCategoryUC)

	getCategoryUC := categoryApp.NewGetCategoryUseCase(cRepo)
	categoryGetController := categoryGetCtrl.NewCategoryGetController(getCategoryUC)

	r := gin.Default()
	r.Use(cors.Default())

	productInfra.SetupProductRoutes(r, pController, productGetController, productBuyController)
	categoryInfra.SetupCategoryRoutes(r, cController, categoryGetController)

	r.Run(":8080")
}
