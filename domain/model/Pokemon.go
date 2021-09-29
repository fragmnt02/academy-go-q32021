package model

import (
	"encoding/json"
)

type Pokemon struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (pokemon *Pokemon) ToJson() ([]byte, error) {
	response, err := json.Marshal(pokemon)
	if err != nil {
		return nil, err
	}

	return response, nil
}
