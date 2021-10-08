package datastore

import (
	"net/http"
)

type API struct {
	url string
}

func (api *API) Init(url string) {
	api.url = url
}

func (api *API) Get(path string) (*http.Response, error) {
	res, err := http.Get(api.url + path)
	return res, err
}
