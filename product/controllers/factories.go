package controllers

import (
	"github.com/JohnAzedo/eCommerce/product/domain"
	"github.com/JohnAzedo/eCommerce/product/infra/database"
	"github.com/JohnAzedo/eCommerce/product/infra/repositories"
)

func ProductControllerFactory() *ProductController {
	return &ProductController{
		domain.NewCreateProductUseCase(
			repositories.ProductRepositoryImpl{
				DB: database.GetInstance(),
			},
			repositories.UUIDRepositoryImpl{},
		),
	}
}
