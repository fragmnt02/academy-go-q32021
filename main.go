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
	db.Init("db.csv")
	repositories := repository.GetRepositories(db)
	controllers := controller.GetControllers(repositories)
	router := router.GetRouter(controllers)
	http.ListenAndServe(":8000", router)

}
