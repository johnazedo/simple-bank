package domain

import "gorm.io/gorm"

type News struct {
	gorm.Model
	ID string `json:"id"`
	Title string `json:"title"`
	Url string `json:"url"`
	Author string `json:"author"`
}