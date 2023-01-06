package src

import (
	"strings"
)

type Server struct {
	ID   string
	Name string
	Host string
	Key  string
	Slug string
}

type RoutesRepository interface {
	GetServer(slug string) (Server, error)
}

type GetServerUseCase struct {
	RoutesRepository
}

func (uc GetServerUseCase) Invoke(path string) (Server, error) {
	slug := strings.Split(strings.TrimPrefix(path, "/"), "/")[0]
	return uc.GetServer(slug)
}
