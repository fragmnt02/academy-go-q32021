package main

import (
	"academy-go-q32021/infrastructure/datastore"
	"academy-go-q32021/infrastructure/router"
	"academy-go-q32021/interface/controller"

	"github.com/gorilla/mux"
)

func main() {
	db := new(datastore.Db)
	db.Init()
	app := controller.InitializeApp(db)
	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	poke_router := r.PathPrefix("/pokemons").Subrouter()
	router.GetPokemonRouter(poke_router, app)
	app.Run(r)

}
