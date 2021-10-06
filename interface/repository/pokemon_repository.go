package repository

import (
	"academy-go-q32021/domain/model"
	"academy-go-q32021/infrastructure/datastore"
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"os"
	"strconv"
)

type PokemonRepository struct {
	db  *datastore.Db
	api *datastore.API
}

func (p *PokemonRepository) Init(db *datastore.Db, api *datastore.API) {
	p.db = db
	p.api = api
}

func (p *PokemonRepository) FindAll() ([]model.Pokemon, error) {
	pokemons := make([]model.Pokemon, len(p.db.Data))
	for i, line := range p.db.Data {
		id, _ := strconv.Atoi(line[0])
		pokemon := model.Pokemon{
			ID:   id,
			Name: line[1],
		}
		pokemons[i] = pokemon
	}
	return pokemons, nil
}

func (p *PokemonRepository) Find(id int) (model.Pokemon, error) {
	var pokemon model.Pokemon
	for _, line := range p.db.Data {
		id_pokemon, _ := strconv.Atoi(line[0])
		if id_pokemon == id {
			pokemon = model.Pokemon{
				ID:   id_pokemon,
				Name: line[1],
			}
		}
	}
	if pokemon.ID == 0 && pokemon.Name == "" {
		return pokemon, errors.New("pokemon not found")
	}
	return pokemon, nil
}

func (p *PokemonRepository) Create(pokemon *model.Pokemon) error {
	id := strconv.Itoa(pokemon.ID)
	var data = make([]string, 2)
	data[0] = id
	data[1] = pokemon.Name
	p.db.Data = append(p.db.Data, data)
	f, err := os.Create("db.csv")
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, record := range p.db.Data {
		if err := w.Write(record); err != nil {
			return err
		}
	}
	return nil
}

func (p *PokemonRepository) FindInAPI(id int) (model.Pokemon, error) {
	var pokemon model.Pokemon
	id_pokemon := strconv.Itoa(id)
	res, err := p.api.Get(id_pokemon)
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
