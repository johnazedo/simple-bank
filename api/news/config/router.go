package config

import (
	"github.com/go-chi/chi"
	"news/controllers"
)

func Router() *chi.Mux{
	router := chi.NewRouter()

	newsController := controllers.NewsController{}
	router.Get("/", newsController.Fetch)
	return router
}