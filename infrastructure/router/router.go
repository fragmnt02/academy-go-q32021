package router

import (
	"net/http"
)

type Router struct {
	Rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		Rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) findHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, pathExist := r.Rules[path]
	handler, methodExist := r.Rules[path][method]
	return handler, methodExist, pathExist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, methodExist, pathExist := r.findHandler(request.URL.Path, request.Method)

	if !pathExist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handler(w, request)
}
