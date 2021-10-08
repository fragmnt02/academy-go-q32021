package tests

import (
	"academy-go-q32021/infrastructure/datastore"
	"academy-go-q32021/interface/controller"
	"academy-go-q32021/interface/repository"
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
	// TODO Check values
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
	// TODO Check values
}
