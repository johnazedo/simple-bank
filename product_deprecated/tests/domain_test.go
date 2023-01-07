package tests

import (
	. "github.com/JohnAzedo/eCommerce/product/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProduct(t *testing.T) {
	product := Product{Price: 0, Amount: -1}
	t.Run("GivenNegativeAmountReturnZero", GivenNegativeAmountReturnZero(product))
	t.Run("GivenPriceZeroReturnError", GivenPriceZeroReturnError(product))
}

func GivenNegativeAmountReturnZero(product Product) func(*testing.T) {
	// Given a product_deprecated with negative amount when call CheckValidAmount then it will return 0
	return func(t *testing.T) {
		product.CheckValidAmount()
		assert.Equal(t, 0, product.Amount, "Product amount should be equal 0")
	}
}

func GivenPriceZeroReturnError(product Product) func(*testing.T) {
	// Given a product_deprecated with price equal or less than 0 when calling CheckValidPrice then it will trigger panic error
	return func(t *testing.T) {
		err := product.CheckValidPrice()
		expected := InvalidPriceError.Error()
		assert.EqualError(t, err, expected, "Error must be InvalidPriceError type")
	}
}
