package tests

import (
	"errors"
	"github.com/JohnAzedo/eCommerce/product/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestProduct(t *testing.T) {
	product := domain.Product{Price: 0, Amount: -1}
	t.Run("GivenNegativeAmountReturnZero", GivenNegativeAmountReturnZero(product))
	t.Run("GivenPriceZeroTriggerPanic", GivenPriceZeroTriggerPanic(product))
}

func GivenNegativeAmountReturnZero(product domain.Product) func(t *testing.T) {
	// Given a product with negative amount when call CheckValidAmount then it will return 0
	return func(t *testing.T) {
		product.CheckValidAmount()
		assert.Equal(t, 0, product.Amount, "Product amount should be equal 0")
	}
}

func GivenPriceZeroTriggerPanic(product domain.Product) func(t *testing.T) {
	// Given a product with price equal or less than 0 when calling CheckValidPrice then it will trigger panic error
	return func(t *testing.T) {
		assert.Panics(t, func() {
			product.CheckValidPrice()
		}, "ValidPrice did not panic")
	}
}


type UUIDRepositoryMock struct {
	mock.Mock
}

func (m *UUIDRepositoryMock) GetNewUUID() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

type ProductRepositoryMock struct {
	mock.Mock
}

func (m *ProductRepositoryMock) CreateProduct(product *domain.Product) error {
	args := m.Called()
	return args.Error(0)
}

func TestCreateProductUseCase(t *testing.T) {
	productRepository := new(ProductRepositoryMock)
	uuidRepository := new(UUIDRepositoryMock)
	useCase := domain.CreateProductUseCase{
		productRepository,
		uuidRepository,
	}

	t.Run(
		"GivenValidProductReturnProduct",
		GivenValidProductReturnSuccessState(useCase, productRepository, uuidRepository))

	t.Run("GivenUUIDErrorReturnErrorState",
		GivenUUIDErrorReturnErrorState(useCase, productRepository, uuidRepository))

	t.Run("GivenCreateProductErrorReturnErrorState",
		GivenCreateProductErrorReturnErrorState(useCase, productRepository, uuidRepository))

}

func GivenValidProductReturnSuccessState(
	useCase domain.CreateProductUseCase,
	productRepository *ProductRepositoryMock,
	uuidRepository *UUIDRepositoryMock) func(t *testing.T) {

	// Given a valid product when calling Execute then it will return ProductCreated
	productRepository.On("CreateProduct").Return(nil)
	uuidRepository.On("GetNewUUID").Return("bb84993a-29be-4fd8-9da4-76037ab37b83", nil)

	return func(t *testing.T) {
		result := useCase.Execute(&domain.Product{Price: 10, Amount: 1})
		assert.Equal(t, domain.ProductCreated, result)
	}
}

// TODO: Fix this
func GivenUUIDErrorReturnErrorState(
	useCase domain.CreateProductUseCase,
	productRepository *ProductRepositoryMock,
	uuidRepository *UUIDRepositoryMock) func(t *testing.T) {

	// Given an error in UUIDRepository when calling GetNewUUID then it will return ProductCreateError
	productRepository.On("CreateProduct").Return(nil)
	uuidRepository.On("GetNewUUID").Return(nil, errors.New("Internal error"))

	return func(t *testing.T) {
		result := useCase.Execute(&domain.Product{Price: 10, Amount: 1})
		assert.Equal(t, domain.ProductCreateError, result)
	}
}

// TODO: Fix this
func GivenCreateProductErrorReturnErrorState(
	useCase domain.CreateProductUseCase,
	productRepository *ProductRepositoryMock,
	uuidRepository *UUIDRepositoryMock) func(t *testing.T) {

	// Given an error when calling CreateProduct then it will return ProductCreateError
	productRepository.On("CreateProduct").Return(errors.New("Internal error"))
	uuidRepository.On("GetNewUUID").Return("bb84993a-29be-4fd8-9da4-76037ab37b83", nil)

	return func(t *testing.T) {
		result := useCase.Execute(&domain.Product{Price: 10, Amount: 1})
		assert.Equal(t, domain.ProductCreateError, result)
	}
}