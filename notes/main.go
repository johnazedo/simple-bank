package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Note struct {
	ID	  string  `json:"id"`;
	Text  string  `json:"text"`;
}

var db []Note;
func GetNotes(response http.ResponseWriter, request *http.Request){
	json.NewEncoder(response).Encode(db)
}

func main() {
	db = append(db, []Note{
		Note{"1", "This is a note"},
		Note{"2", "This is a note to test"},
	}...)

	router := mux.NewRouter()
	router.HandleFunc("/notes", GetNotes).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", router))
}
