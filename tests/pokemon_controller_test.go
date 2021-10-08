package tests

import (
	"academy-go-q32021/domain/model"
	"academy-go-q32021/infrastructure/datastore"
	"academy-go-q32021/interface/controller"
	"academy-go-q32021/interface/repository"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

var mockDB = datastore.Db{
	Data: [][]string{{"1", "test pokemon 1"}, {"2", "test pokemon 2"}},
}

var mockAPI datastore.API

var mockRepository repository.Repositories

var mockController controller.Controllers

func TestHandleGetAllPokemons(t *testing.T) {
	mockRepository = *repository.GetRepositories(&mockDB, &mockAPI)
	mockController = *controller.GetControllers(&mockRepository)

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	mockController.PokemonController.HandleGetAllPokemons(w, req)
	res := w.Result()
	if res.StatusCode != 200 {
		t.Fatalf("No successful request")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}
	var pokemons []model.Pokemon
	err = json.Unmarshal(body, &pokemons)
	if err != nil {
		t.Fatalf("could not unmarshall response %v", err)
	}
	if len(pokemons) != 2 {
		t.Fatalf("response was incomplete")
	}
}

func TestHandleGetPokemon(t *testing.T) {
	mockRepository = *repository.GetRepositories(&mockDB, &mockAPI)
	mockController = *controller.GetControllers(&mockRepository)
	req := httptest.NewRequest("GET", "/1", nil)
	vars := map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)
	w := httptest.NewRecorder()
	mockController.PokemonController.HandleGetPokemon(w, req)
	res := w.Result()
	print(res)
	if res.StatusCode != 200 {
		t.Fatalf("No successful request")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}
	var pokemon model.Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		t.Fatalf("could not unmarshall response %v", err)
	}
	if pokemon.ID != 1 || pokemon.Name != "test pokemon 1" {
		t.Fatalf("pokemon not found")
	}
}
