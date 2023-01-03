package src

import "net/http"

type Route struct {
	ID   string
	Name string
	Host string
	Key  string
	Slug string
}

type SendRequestUseCase struct{}

func (uc *SendRequestUseCase) Invoke(request *http.Request) (*http.Response, error) {
	// To test use this api: https://uselessfacts.jsph.pl/
	return nil, nil
}
