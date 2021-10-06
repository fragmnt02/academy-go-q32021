package repository

import "academy-go-q32021/infrastructure/datastore"

type Repositories struct {
	PokemonRepository PokemonRepository
}

func GetRepositories(db *datastore.Db, api *datastore.API) *Repositories {
	repositories := new(Repositories)
	pokemonRepository := new(PokemonRepository)
	pokemonRepository.Init(db, api)
	repositories.PokemonRepository = *pokemonRepository
	return repositories
}
