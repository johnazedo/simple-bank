package controllers

import "github.com/gin-gonic/gin"

func SetupServer(controller *ProductController) *gin.Engine {
	router := gin.Default()
	api := router.Group("/products")
	api.POST("", controller.CreateProduct)
	return router
}