package repositories

import (
	. "github.com/JohnAzedo/eCommerce/product/domain"
	. "github.com/JohnAzedo/eCommerce/product/infra/database"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	ProductRepository
	DB *gorm.DB
}

func (repo ProductRepositoryImpl) CreateProduct(product *Product) error {
	var model *ProductModel
	model = FromProductEntity(product)
	if err := repo.DB.Create(model).Error; err != nil { return err }
	return nil
}