package controller

import (
	"academy-go-q32021/infrastructure/datastore"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleGetAllPokemons(w http.ResponseWriter, r *http.Request) {
	db := new(datastore.Db)
	db.Init()
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
