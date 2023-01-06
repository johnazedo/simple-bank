package src

import (
	"errors"
	"gorm.io/gorm"
)

type RoutesRepositoryImpl struct {
	db *gorm.DB
}

func (r *RoutesRepositoryImpl) GetServer(slug string) (Server, error) {
	var route Server
	err := r.db.First(&route, "slug = ?", slug).Error
	return route, err
}

type RoutesRepositoryFake struct{}

func (r RoutesRepositoryFake) GetServer(slug string) (Server, error) {
	if slug == "useless" {
		return Server{
			ID:   "1",
			Name: "Random",
			Host: "https://uselessfacts.jsph.pl/",
			Slug: "random",
		}, nil
	} else {
		return Server{}, errors.New("route not found")
	}
}
