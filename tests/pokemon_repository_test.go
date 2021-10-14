package tests

import (
	"academy-go-q32021/interface/repository"
	"testing"
)

func TestFind(t *testing.T) {
	mockRepository = *repository.GetRepositories(&mockDB)
	pokemon, err := mockRepository.PokemonRepository.Find(1)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if pokemon.Name != "test pokemon 1" || pokemon.ID != 1 {
		t.Fatalf("pokemon not found correctly")
	}
}

func TestFindAll(t *testing.T) {
	mockRepository = *repository.GetRepositories(&mockDB)
	pokemons, err := mockRepository.PokemonRepository.FindAll()
	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(pokemons) != 2 {
		t.Fatalf("Not the correct amount of pokemons")
	}
}

func TestFindInAPI(t *testing.T) {
	mockRepository = *repository.GetRepositories(&mockDB)
	pokemon, err := mockRepository.PokemonAPIRepository.Find(1)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if pokemon.Name != "bulbasaur" || pokemon.ID != 1 {
		t.Fatalf("pokemon not found correctly")
	}
}
