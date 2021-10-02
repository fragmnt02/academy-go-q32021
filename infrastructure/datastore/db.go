package datastore

import (
	"academy-go-q32021/domain/model"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

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
