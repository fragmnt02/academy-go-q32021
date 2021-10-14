package router

import (
	"academy-go-q32021/interface/controller"

	"github.com/gorilla/mux"
)

// GetRouter(controllers *controller.Controllers): Return the router with all routers already set.
func GetRouter(controllers *controller.Controllers) *mux.Router {
	mainRouter := mux.NewRouter().PathPrefix("/api").Subrouter()
	pokeRouter := mainRouter.PathPrefix("/pokemons").Subrouter()
	getPokemonRouter(pokeRouter, controllers)
	return mainRouter
}
