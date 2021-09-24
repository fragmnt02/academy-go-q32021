package main

import (
	"net/http" 
	"fmt"
	"encoding/json"
)

func handleGetAllPokemons(w http.ResponseWriter, r *http.Request)  {
	db := new(Db)
	db.Init()
	pokemons, err := db.GetAllPokemons()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	response, err2 := json.Marshal(pokemons)
    if err2 != nil {
        w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err2.Error())
		return
    }
	w.Header().Set("Content-Type","application/json")
	w.Write(response)
}