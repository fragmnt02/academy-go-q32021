package controller

import (
	"academy-go-q32021/interface/repository"
)

type Controllers struct {
	PokemonController PokemonController
}

func GetControllers(r *repository.Repositories) *Controllers {
	controllers := new(Controllers)
	pokemonController := new(PokemonController)
	pokemonController.init(r)
	controllers.PokemonController = *pokemonController
	return controllers
}
