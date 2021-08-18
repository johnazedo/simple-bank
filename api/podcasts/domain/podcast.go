package domain


type Podcast struct {
	Base
	Title string `json:"title"`
	Number uint8
	AudioUrl string
}

// Before created we must save audio in AWS server

