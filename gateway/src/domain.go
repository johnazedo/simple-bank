package src

import "net/http"

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

type RequestRepository interface {
	SendRequest(response http.Response) (*http.Response, error)
}

type SendRequestUseCase struct {
	RoutesRepository
	RequestRepository
}

func (uc *SendRequestUseCase) Invoke(request *http.Request) (*http.Response, error) {
	// To test use this api: https://uselessfacts.jsph.pl/random.json?language=en
	return nil, nil
}
