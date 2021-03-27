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
func GetNotes(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(db)
}

func CreateNotes(w http.ResponseWriter, r *http.Request){
	var note *Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db = append(db, *note)
	json.NewEncoder(w).Encode(*note)
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/notes", GetNotes).Methods("GET")
	router.HandleFunc("/notes", CreateNotes).Methods("POST")
	return router
}

func main() {
	db = append(db, []Note{
		Note{"1", "This is a note"},
		Note{"2", "This is a note to test"},
	}...)

	log.Fatal(http.ListenAndServe(":10000", Router()))
}
