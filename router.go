package main

import(
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func newRouter() *Router  {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) findHandler(path string,method string) (http.HandlerFunc, bool, bool)  {
	_, pathExist := r.rules[path]
	handler, methodExist:= r.rules[path][method]
	return handler, methodExist, pathExist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request)  {
	handler, methodExist, pathExist := r.findHandler(request.URL.Path, request.Method)

	if !pathExist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handler(w,request)
}