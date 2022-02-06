package infra

import (
	"errors"
	. "github.com/JohnAzedo/eCommerce/product/domain"
)

type ProductRepositoryImpl struct {
	ProductRepository
}

func (repo ProductRepositoryImpl) CreateProduct(product *Product) error {
	return errors.New("Not implementated")
}