package main

import (
	"encoding/csv"
    "fmt"
    "os"
	"strconv"
)

type Db struct {
	data [][]string
}

func (db *Db) Init() error {
	csvFile, err1 := os.Open("db.csv")
	if err1 != nil {
		return err1
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()
    
    csvLines, err2 := csv.NewReader(csvFile).ReadAll()
    if err2 != nil {
        return err2
    } 
	db.data = csvLines
	return nil
}

func (db *Db) GetAllPokemons() ([]Pokemon, error) {
	pokemons := make([]Pokemon,len(db.data))
	for i, line := range db.data {
        id, _ := strconv.Atoi(line[0])
		pokemon := Pokemon {
			ID: id,
			Name: line[1],
		}
		pokemons[i] = pokemon
    }
	return pokemons, nil
}