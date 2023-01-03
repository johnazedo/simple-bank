package main

import (
	. "github.com/johnazedo/marketplace/gateway/src"
	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()
	controller := RouterController{
		UseCase: SendRequestUseCase{},
	}
	server.GET("/:", controller.Get)
	server.Logger.Fatal(server.Start("0.0.0.0:8000"))
}
