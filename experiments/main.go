package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

var users []User

func GetUsers(context echo.Context) error {
	return context.JSON(http.StatusOK, users)
}

func LoadUsers() {
	users = append(users, User{
		Name:     "John",
		Username: "Lemon",
	})
}

// Create a webserver using Echo
func main() {
	LoadUsers()
	e := echo.New()
	e.GET("/", HelloWorld)
	e.GET("/users", GetUsers)
	e.Logger.Fatal(e.Start(":8000"))
}

func HelloWorld(context echo.Context) error {
	return context.JSON(http.StatusOK, "Hello, World!")
}
