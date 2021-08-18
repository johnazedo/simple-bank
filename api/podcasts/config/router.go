package config

import (
	"github.com/go-chi/chi"
)

func Router() *chi.Mux {
	router := chi.NewRouter();
	return router;
}