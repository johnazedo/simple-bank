package main

import (
	"github.com/go-chi/chi/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	_ "net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-chi/chi/middleware"
)

func main(){
	_, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}


	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", nil)
	_ = http.ListenAndServe(":3000", router)
}