package main

import (
	"github.com/JohnAzedo/eCommerce/product/controllers"
	"github.com/JohnAzedo/eCommerce/product/infra/database"
	"github.com/gin-gonic/gin"
	_ "net/http"
)

func main(){
	db := database.GetInstance()
	database.Migrate(db)

	router := gin.Default()

	productController := controllers.ProductControllerFactory()
	api := router.Group("/products")
	api.POST("", productController.CreateProduct)

	_ = router.Run(":3000")
}