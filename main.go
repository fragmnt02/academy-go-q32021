package main

import "academy-go-q32021/interface/controller"

func main() {
	server := controller.NewServer(":3000")
	server.Handle("GET", "/pokemons", controller.HandleGetAllPokemons)
	server.Listen()
}
