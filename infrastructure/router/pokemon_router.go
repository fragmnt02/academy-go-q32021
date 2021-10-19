package router

import (
	"academy-go-q32021/interface/controller"

	"github.com/gorilla/mux"
)

func getPokemonRouter(r *mux.Router, controllers *controller.Controllers) {
	r.HandleFunc("", controllers.PokemonController.HandleGetAllPokemons)
	r.HandleFunc("/", controllers.PokemonController.HandleGetAllPokemons)
	r.HandleFunc("/concurrency", controllers.PokemonController.HandleGetAllPokemonsConcurrently)
	r.HandleFunc("/{id}", controllers.PokemonController.HandleGetPokemon)
}
