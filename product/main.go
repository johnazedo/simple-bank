package main

import (
	"github.com/JohnAzedo/eCommerce/product/infra/database"
	_ "github.com/go-chi/chi/middleware"
	_ "net/http"
)

func main(){
	db := database.GetInstance()
	database.Migrate(db)

	//router := chi.NewRouter()
	//router.Use(middleware.Logger)
	//router.Get("/", nil)
	//_ = http.ListenAndServe(":3000", router)
}