package main

import (
	"net/http"
	"news/config"
	"news/database"
)


func main(){
	database.DatabaseSetup()
	http.ListenAndServe(":8001", config.Router())
}

