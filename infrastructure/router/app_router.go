package router

import (
	"academy-go-q32021/interface/controller"

	"github.com/gorilla/mux"
)

func GetRouter(controllers *controller.Controllers) *mux.Router {
	mainRouter := mux.NewRouter().PathPrefix("/api").Subrouter()
	pokeRouter := mainRouter.PathPrefix("/pokemons").Subrouter()
	getPokemonRouter(pokeRouter, controllers)
	return mainRouter
}
