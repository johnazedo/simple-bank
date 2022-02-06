package main

import (
	"github.com/JohnAzedo/eCommerce/product/infra/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-chi/chi/middleware"
	"net/http"
	_ "net/http"
)

func main(){
	db := database.GetInstance()
	database.Migrate(db)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", nil)
	_ = http.ListenAndServe(":3000", router)
}