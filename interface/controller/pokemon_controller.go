package controller

import (
	"academy-go-q32021/infrastructure/datastore"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HandleGetAllPokemons(db *datastore.Db, w http.ResponseWriter, r *http.Request) {
	pokemons, err := db.GetAllPokemons()
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

func HandleGetPokemon(db *datastore.Db, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id_str := params["id"]
	if id_str == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request")
		return
	}
	id, _ := strconv.Atoi(id_str)
	pokemon, err := db.GetPokemon(id)
	if err != nil {
		if err.Error() == "pokemon not found" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprint(w, err.Error())
		return
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
