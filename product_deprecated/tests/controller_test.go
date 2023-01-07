package tests

import (
	"bytes"
	"encoding/json"
	"github.com/JohnAzedo/eCommerce/product/controllers"
	"github.com/JohnAzedo/eCommerce/product/domain"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProductController(t *testing.T) {
	t.Run("TestInvalidReqBody", GivenInvalidReqBodyWhenCreateProductReturnStatusBadRequest())
}

func GivenInvalidReqBodyWhenCreateProductReturnStatusBadRequest() func(*testing.T){
	controller := controllers.ProductController{CreateProductUseCaseMock{},}
	reqBody, _ := json.Marshal(map[string]interface{}{
		"name":"",
		"image":"",
		"measurement":"",
		"amount":0,
	})

	return func(t *testing.T){
		ts := httptest.NewServer(controllers.SetupServer(&controller))
		defer ts.Close()
		response, _ := http.Post(ts.URL + "/products/create",  "application/json", bytes.NewBuffer(reqBody))
		assert.Equal(t, response.StatusCode, http.StatusBadRequest)
	}
}

type CreateProductUseCaseMock struct {
	domain.CreateProductUseCase
}

func (uc CreateProductUseCaseMock) Execute(product *domain.Product) error {
	return nil
}
