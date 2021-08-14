package domain

type News struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
	Url string `json:"url"`
	Author string `json:"author"`
}