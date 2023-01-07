package main

import (
	"github.com/JohnAzedo/eCommerce/product/controllers"
	"github.com/JohnAzedo/eCommerce/product/infra/postgres"
	_ "net/http"
)

func main(){
	db := postgres.GetInstance()
	postgres.Migrate(db)
	controller := controllers.ProductControllerFactory()
	_ = controllers.SetupServer(controller).Run(":3000")
}

