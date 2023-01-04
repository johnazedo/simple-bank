package src

import (
	"gorm.io/gorm"
	"net/http"
)

type RoutesRepositoryImpl struct {
	db *gorm.DB
}

func (r *RoutesRepositoryImpl) GetRoute(slug string) (Route, error) {
	var route Route
	err := r.db.First(&route, "slug = ?", slug).Error
	return route, err
}

type RoutesRepositoryFake struct{}

func (r *RoutesRepositoryFake) GetRoute(slug string) (Route, error) {
	return Route{
		ID:   "1",
		Name: "Random",
		Host: "https://uselessfacts.jsph.pl/",
		Slug: "random",
	}, nil
}

type RequestRepositoryImpl struct{}

func (r *RequestRepositoryImpl) SendRequest(response http.Response) (*http.Response, error) {
	return nil, nil
}
