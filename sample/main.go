package main

import "sample/chttp"

func main(){
	server := chttp.CreateServer("192.168.0.40", "8080")
	server.Run()
}

