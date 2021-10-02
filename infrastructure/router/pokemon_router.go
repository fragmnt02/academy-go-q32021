package router

import (
	"academy-go-q32021/interface/controller"

	"github.com/gorilla/mux"
)

func GetPokemonRouter(r *mux.Router, app *controller.App) {
	r.HandleFunc("", app.HandleRequest(controller.HandleGetAllPokemons))
	r.HandleFunc("/", app.HandleRequest(controller.HandleGetAllPokemons))
	r.HandleFunc("/{id}", app.HandleRequest(controller.HandleGetPokemon))
}
