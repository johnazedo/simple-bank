package controllers

import (
	"github.com/JohnAzedo/eCommerce/product/domain"
	"github.com/JohnAzedo/eCommerce/product/infra"
	"github.com/JohnAzedo/eCommerce/product/infra/postgres"
)

func ProductControllerFactory() *ProductController {
	return &ProductController{
		domain.NewCreateProductUseCase(
			postgres.ProductRepositoryImpl{
				DB: postgres.GetInstance(),
			},
			infra.UUIDRepositoryImpl{},
		),
	}
}
