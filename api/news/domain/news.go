package domain

type News struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Url string `json:"url"`
	Author string `json:"author"`
}