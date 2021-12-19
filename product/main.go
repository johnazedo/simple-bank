package main

import "net/http"

func main(){
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return 
	}
}