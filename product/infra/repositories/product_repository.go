package repositories

import (
	. "github.com/JohnAzedo/eCommerce/product/domain"
	. "github.com/JohnAzedo/eCommerce/product/infra/database"
)

type ProductRepositoryImpl struct {
	ProductRepository
}

func (repo ProductRepositoryImpl) CreateProduct(product *Product) error {
	var model *ProductModel
	model = FromProductEntity(product)
	if err := GetInstance().Create(model).Error; err != nil { return err }
	return nil
}