package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"news/database"
	"news/domain"
)

type NewsController struct{}

func (nc *NewsController) FetchAll(w http.ResponseWriter, r *http.Request) {
	var news []domain.News
	database.GetDB().Find(&news)
	json.NewEncoder(w).Encode(&news)
}

func (nc *NewsController) Create(w http.ResponseWriter, r *http.Request) {
	var news domain.News;
	err := json.NewDecoder(r.Body).Decode(&news)

	if err != nil {
		log.Fatalln(err)
	}

	database.GetDB().Create(&news)
	json.NewEncoder(w).Encode(&news)
}


func (nc *NewsController) Get(w http.ResponseWriter, r *http.Request) {
	var news domain.News;
	pk := chi.URLParam(r, "pk")
	database.GetDB().First(&news, pk)
	json.NewEncoder(w).Encode(&news)
}