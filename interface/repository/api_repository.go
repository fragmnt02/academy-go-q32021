package repository

import "net/http"

type APIRepository struct {
	url string
}

func (api *APIRepository) init(url string) {
	api.url = url
}

// get(path string): make an http get request
func (api *APIRepository) get(path string) (*http.Response, error) {
	res, err := http.Get(api.url + path)
	return res, err
}
