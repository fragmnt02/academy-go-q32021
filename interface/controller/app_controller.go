package controller

import (
	"academy-go-q32021/infrastructure/router"
	"net/http"
)

type Server struct {
	port   string
	router *router.Router
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: router.NewRouter(),
	}
}

func (s *Server) Handle(method string, path string, handler http.HandlerFunc) {
	_, exist := s.router.Rules[path]
	if !exist {
		s.router.Rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.Rules[path][method] = handler
}

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}
