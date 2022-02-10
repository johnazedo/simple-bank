package controllers

import (
	"errors"
	"fmt"
	"github.com/JohnAzedo/eCommerce/product/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductController struct {
	*domain.CreateProductUseCase
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	usecase := pc.CreateProductUseCase
	product := domain.Product{}
	err := c.Bind(&product)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":"failed", "message": errors.New("Invalid request body")})
		return
	}

	err = usecase.Execute(&product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": errors.New("Internal error")})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status":"success", "user": &product})

}