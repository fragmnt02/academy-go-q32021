package repository

import (
	"academy-go-q32021/domain/model"
	"academy-go-q32021/infrastructure/datastore"
	"errors"
	"math"
	"strconv"
)

type PokemonRepository struct {
	db *datastore.Db
}

func (p *PokemonRepository) init(db *datastore.Db) {
	p.db = db
}

// FindAll():  Return the list of all pokemons in the csv database
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

// Find(id int): Get the pokemon with the given id in the csv database
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

// Create(Pokemon model.Pokemon): Save a new pokemon in the csv database
func (p *PokemonRepository) Create(pokemon *model.Pokemon) error {
	id := strconv.Itoa(pokemon.ID)
	data := make([]string, 2)
	data[0] = id
	data[1] = pokemon.Name
	p.db.WriteLine(data)
	return nil
}

func (p *PokemonRepository) FindAllConcurrently(items int, itemsPerWorker int, typeValue string) []model.Pokemon {
	jobs := make(chan []string, items)
	results := make(chan []string, items)
	workers := math.Ceil(float64(items) / float64(itemsPerWorker))
	for i := 0; i <= int(workers); i++ {
		go worker(jobs, results, typeValue)
	}

	for index, line := range p.db.Data {
		jobs <- line
		if index >= items {
			break
		}
	}

	close(jobs)

	var result = []model.Pokemon{}

	for a := 1; a <= items; a++ {
		data := <-results
		if data != nil {
			id, _ := strconv.Atoi(data[0])
			result = append(result, model.Pokemon{
				ID:   id,
				Name: data[1],
			})
		}
	}

	return result
}

func worker(jobs <-chan []string, results chan<- []string, divisibility string) {
	for job := range jobs {
		results <- getItem(job, divisibility)
	}
}

func getItem(item []string, div string) []string {
	id, _ := strconv.Atoi(item[0])
	var isEven = id%2 == 0
	if (div == "odd" && !isEven) || (div == "even" && isEven) {
		return item
	}
	return nil
}
