package repository

import "net/http"

type APIRepository struct {
	url string
}

func (api *APIRepository) Init(url string) {
	api.url = url
}

func (api *APIRepository) Get(path string) (*http.Response, error) {
	res, err := http.Get(api.url + path)
	return res, err
}
