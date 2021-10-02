package controller

import (
	"academy-go-q32021/infrastructure/datastore"
	"encoding/json"
	"fmt"
	"net/http"
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
