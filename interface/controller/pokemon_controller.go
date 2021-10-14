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

func (p *PokemonController) initialize(pr *repository.Repositories) {
	p.pr = pr
}

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
