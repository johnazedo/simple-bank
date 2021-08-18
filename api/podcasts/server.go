package main

import (
	"net/http"
	"podcasts/config"
)

func main() {
	http.ListenAndServe(":8002", config.Router())
}


