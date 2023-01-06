package src

import (
	"strings"
)

type Route struct {
	ID   string
	Name string
	Host string
	Key  string
	Slug string
}

type RoutesRepository interface {
	GetRoute(slug string) (Route, error)
}

type GetRouteUseCase struct {
	RoutesRepository
}

func (uc *GetRouteUseCase) Invoke(path string) (Route, error) {
	// To test use this api: https://uselessfacts.jsph.pl/random.json?language=en
	slug := strings.SplitAfter(path, "/")[0]
	return uc.GetRoute(slug)
}
