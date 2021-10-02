package datastore

import (
	"academy-go-q32021/domain/model"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

const POKE_API_URL = "https://pokeapi.co/api/v2/pokemon/"

type Db struct {
	data [][]string
}

func (db *Db) Init() error {
	csvFile, err := os.Open("db.csv")
	if err != nil {
		return err
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return err
	}
	db.data = csvLines
	return nil
}

func (db *Db) GetAllPokemons() ([]model.Pokemon, error) {
	pokemons := make([]model.Pokemon, len(db.data))
	for i, line := range db.data {
		id, _ := strconv.Atoi(line[0])
		pokemon := model.Pokemon{
			ID:   id,
			Name: line[1],
		}
		pokemons[i] = pokemon
	}
	return pokemons, nil
}

func (db *Db) GetPokemon(id_pokemon int) (model.Pokemon, error) {
	var pokemon model.Pokemon
	for _, line := range db.data {
		id, _ := strconv.Atoi(line[0])
		if id_pokemon == id {
			pokemon = model.Pokemon{
				ID:   id,
				Name: line[1],
			}
		}
	}
	if pokemon.ID == 0 && pokemon.Name == "" {
		return pokemon, errors.New("pokemon not found")
	}
	return pokemon, nil
}

func (db *Db) GetPokemonFromApi(id_pokemon int) (model.Pokemon, error) {
	var pokemon model.Pokemon
	id := strconv.Itoa(id_pokemon)
	res, err := http.Get(POKE_API_URL + id)
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

func (db *Db) SavePokemon(pokemon *model.Pokemon) error {
	id := strconv.Itoa(pokemon.ID)
	var data = make([]string, 2)
	data[0] = id
	data[1] = pokemon.Name
	db.data = append(db.data, data)
	f, err := os.Create("db.csv")
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, record := range db.data {
		if err := w.Write(record); err != nil {
			return err
		}
	}
	return nil
}
