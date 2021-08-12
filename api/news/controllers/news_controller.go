package controllers

import (
	"encoding/json"
	"net/http"
	"news/database"
)

type NewsController struct{}

func (nc *NewsController) Fetch(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(database.GetAllNews())
}
