package tests

import (
	"context"
	"github.com/JohnAzedo/eCommerce/product/domain"
	"github.com/JohnAzedo/eCommerce/product/infra/postgres"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductRepository(t *testing.T){
	ctx := context.Background()
	container, err := SetupContainer(ctx)
	if err != nil {
		panic(err)
	}
	repository := postgres.ProductRepositoryImpl{DB: container.Database}
	t.Run("GivenValidProductReturnSuccess", GivenValidProductReturnSuccess(repository))
	container.Terminate(ctx)
}

func GivenValidProductReturnSuccess(repository domain.ProductRepository) func(*testing.T){
	product := &domain.Product{
		UUID: "24c0bbcd-83ab-43f2-b1a4-ebfe74d3fe03",
		Name: "Shoes",
		Image: "image_url",
		Price: 60.0,
		Description: "Is a shoes",
		Measurement: domain.PerUnit,
		Amount: 20,
	}
	return func(t *testing.T) {
		err := repository.CreateProduct(product)
		assert.Equal(t, nil, err)
	}
}