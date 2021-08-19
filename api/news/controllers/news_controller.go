package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"io"
	"log"
	"net/http"
	"news/database"
	"news/domain"
	"os"
	"strings"
)

type NewsController struct{}

func (nc *NewsController) FetchAll(w http.ResponseWriter, r *http.Request) {
	var news []domain.News
	database.GetDB().Find(&news)
	json.NewEncoder(w).Encode(&news)
}

func (nc *NewsController) Create(w http.ResponseWriter, r *http.Request) {
	var news domain.News;

	uploadFile(r)
	err := json.NewDecoder(r.Body).Decode(&news)

	if err != nil {
		log.Fatalln(err)
	}

	database.GetDB().Create(&news)
	json.NewEncoder(w).Encode(&news)
}

func uploadFile(r *http.Request) (*os.File) {
	err := r.ParseMultipartForm(23 << 20);
	if err != nil {
		log.Println("PARSE_ERROR: Cannot get file!")
		return nil
	}

	image, header, err := r.FormFile("imageUrl");
	defer image.Close();

	if err != nil {
		log.Println("KEY_ERROR: Cannot get file!")
		return nil
	}

	name := strings.Split(header.Filename, ".");
	log.Println("File name ", name[0]);

	file, err := os.OpenFile("./downloaded", os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()

	if err != nil {
		log.Fatalln("OS_FILE_ERROR: Cannot open os file!")
	}

	io.Copy(file, image)
	return file
}


func (nc *NewsController) Get(w http.ResponseWriter, r *http.Request) {
	var news domain.News;
	pk := chi.URLParam(r, "pk")
	database.GetDB().First(&news, pk)
	json.NewEncoder(w).Encode(&news)
}

func (nc *NewsController) Delete(w http.ResponseWriter, r *http.Request) {
	var news domain.News;
	result := make(map[string]string)
	result["msg"] = "Object deleted!"

	pk := chi.URLParam(r, "pk")

	err := database.GetDB().Delete(&news, pk).Error
	if err != nil {
		log.Println("Cannot delete this!")
		result["msg"] = "Invalid Request"
	}

	json.NewEncoder(w).Encode(result)
}
