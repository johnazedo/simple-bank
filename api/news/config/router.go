package config

import (
	"github.com/go-chi/chi"
	"news/controllers"
)

func Router() *chi.Mux{
	router := chi.NewRouter()

	controller := controllers.NewsController{}
	router.Get("/",  controller.FetchAll)
	router.Post("/create", controller.Create)
	router.Get("/{pk}", controller.Get)
	router.Delete("/{pk}/delete", controller.Delete)
	return router
}