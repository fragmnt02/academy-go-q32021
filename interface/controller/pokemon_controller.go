package controller

import (
	"academy-go-q32021/interface/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PokemonController struct {
	pr *repository.Repositories
}

func (p *PokemonController) init(pr *repository.Repositories) {
	p.pr = pr
}

// HandleGetAllPokemons(w http.ResponseWriter, r *http.Request): Controller to get all pokemons in database
func (p *PokemonController) HandleGetAllPokemons(w http.ResponseWriter, r *http.Request) {
	pokemons, err := p.pr.PokemonRepository.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	response, err := json.Marshal(pokemons)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// HandleGetPokemon(w http.ResponseWriter, r *http.Request): Controller to get a specific pokemon, should have an id value in params
func (p *PokemonController) HandleGetPokemon(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id_str := params["id"]
	if id_str == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request")
		return
	}
	id, _ := strconv.Atoi(id_str)
	pokemon, err := p.pr.PokemonRepository.Find(id)

	if err != nil {
		if err.Error() == "pokemon not found" {
			pokemonApi, err := p.pr.PokemonAPIRepository.Find(id)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			} else {
				p.pr.PokemonRepository.Create(&pokemonApi)
				pokemon = pokemonApi
			}
		} else {
			fmt.Fprint(w, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	response, err := json.Marshal(pokemon)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// HandleGetAllPokemons(w http.ResponseWriter, r *http.Request): Controller to get all pokemons in database
func (p *PokemonController) HandleGetAllPokemonsConcurrently(w http.ResponseWriter, r *http.Request) {
	type_string := r.URL.Query().Get("type")
	items_string := r.URL.Query().Get("items")
	items_per_workers_string := r.URL.Query().Get("items_per_workers")

	if type_string == "" || items_string == "" || items_per_workers_string == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "You need to send type, items and items_per_worker in query parameters.")
		return
	}

	if type_string != "odd" && type_string != "even" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "type query parameter should be \"even\" or \"odd\".")
		return
	}

	items_per_workers, err := strconv.Atoi(items_per_workers_string)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "items_per_workers needs to be an integer.")
		return
	}

	items, err := strconv.Atoi(items_string)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "items needs to be an integer.")
		return
	}

	if items_per_workers > items {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "items_per_workers should be less than item.")
		return
	}

	pokemons := p.pr.PokemonRepository.FindAllConcurrently(items, items_per_workers, type_string)

	response, err := json.Marshal(pokemons)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
