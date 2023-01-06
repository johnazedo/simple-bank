package main

import (
	"github.com/gin-gonic/gin"
	. "github.com/johnazedo/marketplace/gateway/src"
	"log"
)

func main() {
	server := gin.Default()
	controller := RouterController{
		GetServerUseCase: GetServerUseCase{
			RoutesRepository: RoutesRepositoryFake{},
		},
	}
	server.Any("/*slug", controller.Proxy)
	if err := server.Run("0.0.0.0:8000"); err != nil {
		log.Fatal(err)
	}
}
