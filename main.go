package main

import (
	"github.com/wtwingate/pokedex/internal/pokeapi"
)

type Config struct {
	Client   pokeapi.Client
	Next     *string
	Previous *string
}

func main() {
	config := Config{
		Client:   *pokeapi.NewClient(),
		Next:     nil,
		Previous: nil,
	}

	startREPL(&config)
}
