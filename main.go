package main

func main() {
    server := newServer(":3000")
    server.Handle("GET","/pokemons",handleGetAllPokemons)
    server.Listen()
}
