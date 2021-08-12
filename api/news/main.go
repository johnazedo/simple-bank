package main

import (
	"net/http"
	"news/config"
)


func main(){
	http.ListenAndServe(":8001", config.Router())
}

