package repository

import "academy-go-q32021/infrastructure/datastore"

type Repositories struct {
	PokemonRepository    PokemonRepository
	PokemonAPIRepository PokemonAPIRepository
}

// GetRepositories(db datastore.Db): Return all the repositories of the app
func GetRepositories(db *datastore.Db) *Repositories {
	pokemonAPIRepository := new(PokemonAPIRepository)
	pokemonAPIRepository.init()

	pokemonRepository := new(PokemonRepository)
	pokemonRepository.init(db)

	repositories := new(Repositories)
	repositories.PokemonRepository = *pokemonRepository
	repositories.PokemonAPIRepository = *pokemonAPIRepository
	return repositories
}
