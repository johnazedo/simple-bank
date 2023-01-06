package main

import (
	"github.com/gin-gonic/gin"
	. "github.com/johnazedo/marketplace/gateway/src"
	"log"
)

func main() {
	server := gin.Default()
	controller := RouterController{
		UseCase: GetRouteUseCase{
			RoutesRepository: RoutesRepositoryFake{},
		},
	}
	server.GET("/:", controller.Proxy)
	if err := server.Run("0.0.0.0:8000"); err != nil {
		log.Fatal(err)
	}
}
