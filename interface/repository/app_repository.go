package repository

import "academy-go-q32021/infrastructure/datastore"

type Repositories struct {
	PokemonRepository    PokemonRepository
	PokemonAPIRepository PokemonAPIRepository
}

func GetRepositories(db *datastore.Db) *Repositories {
	repositories := new(Repositories)
	pokemonAPIRepository := new(PokemonAPIRepository)
	pokemonAPIRepository.Init()
	pokemonRepository := new(PokemonRepository)
	pokemonRepository.Init(db)
	repositories.PokemonRepository = *pokemonRepository
	repositories.PokemonAPIRepository = *pokemonAPIRepository
	return repositories
}
