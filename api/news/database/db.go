package database

import (
	"news/domain"
)

var db[]domain.News;

func GetAllNews() []domain.News {
	db = append(db, []domain.News{
		{"1", "This is a test", "http://newsletter.com", "johnazedo"},
	}...)

	return db
}
