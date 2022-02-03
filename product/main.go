package main

import (
	"github.com/go-chi/chi/middleware"
	"net/http"
	_ "net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-chi/chi/middleware"
)

func main(){
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", nil)
	_ = http.ListenAndServe(":3000", router)
}