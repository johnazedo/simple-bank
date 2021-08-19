package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type News struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Title string `json:"title"`
	Url string `json:"url"`
	Author string `json:"author"`
	Image string `json:"imageUrl"`
}


func (news* News) BeforeCreate(*gorm.DB) error{
	news.ID = uuid.New()
	return nil
}