package src

import (
	"errors"
	"gorm.io/gorm"
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

func (r RoutesRepositoryFake) GetRoute(slug string) (Route, error) {
	if slug == "random.json" {
		return Route{
			ID:   "1",
			Name: "Random",
			Host: "https://uselessfacts.jsph.pl/",
			Slug: "random",
		}, nil
	} else {
		return Route{}, errors.New("route not found")
	}
}
