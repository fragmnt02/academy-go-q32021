package main

import (
	"academy-go-q32021/infrastructure/datastore"
	"academy-go-q32021/infrastructure/router"
	"academy-go-q32021/interface/controller"
	"academy-go-q32021/interface/repository"
	"net/http"
)

func main() {
	db := new(datastore.Db)
	db.Init()
	api := new(datastore.API)
	api.Init("https://pokeapi.co/api/v2/pokemon/")
	repositories := repository.GetRepositories(db, api)
	controllers := controller.GetControllers(repositories)
	router := router.GetRouter(controllers)
	http.ListenAndServe(":8000", router)

}
