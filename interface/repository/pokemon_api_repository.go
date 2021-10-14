package repository

import (
	"academy-go-q32021/domain/model"
	"encoding/json"
	"io"
	"strconv"
)

type PokemonAPIRepository struct {
	apiRepository APIRepository
}

func (par *PokemonAPIRepository) Init() {
	apiRepository := new(APIRepository)
	apiRepository.Init("https://pokeapi.co/api/v2/pokemon/")
	par.apiRepository = *apiRepository
}

func (par *PokemonAPIRepository) Find(id int) (model.Pokemon, error) {
	var pokemon model.Pokemon
	id_pokemon := strconv.Itoa(id)
	res, err := par.apiRepository.Get(id_pokemon)
	if err != nil {
		return pokemon, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return pokemon, err
	}
	var pokemonData = &model.Pokemon{}
	json.Unmarshal([]byte(string(body)), &pokemonData)
	pokemon.ID = pokemonData.ID
	pokemon.Name = pokemonData.Name
	return pokemon, nil
}
