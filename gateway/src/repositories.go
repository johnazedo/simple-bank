package src

import (
	"gorm.io/gorm"
	"net/http"
)

type RoutesRepository struct {
	db *gorm.DB
}

func (r *RoutesRepository) GetRoute(slug string) (Route, error) {
	var route Route
	err := r.db.First(&route, "slug = ?", slug).Error
	return route, err
}

type RequestRepository struct {
}

func (r *RequestRepository) SendRequest(response http.Response) (*http.Response, error) {
	return nil, nil
}