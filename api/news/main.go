package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)


type News struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Url string `json:"url"`
	Author string `json:"author"`
}

var db []News;

func FetchNews(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(db)
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", FetchNews).Methods("GET")
	return router
}

func main(){
	startDB()
	fmt.Println("Server started!")
	http.ListenAndServe(":8001", Router())

}

func startDB(){
	db = append(db, []News{
		{"1", "This is a test", "http://newsletter.com", "johnazedo"},
	}...)
}