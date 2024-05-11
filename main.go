package main

import (
	"github.com/wtwingate/pokedex/internal/pokeapi"
)

type Config struct {
	Client   pokeapi.Client
	Pokedex  map[string]pokeapi.PokemonID
	Next     *string
	Previous *string
}

func main() {
	pokedex := make(map[string]pokeapi.PokemonID)
	config := Config{
		Client:   *pokeapi.NewClient(),
		Pokedex:  pokedex,
		Next:     nil,
		Previous: nil,
	}

	startREPL(&config)
}
