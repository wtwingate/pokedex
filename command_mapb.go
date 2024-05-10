package main

import (
	"errors"
	"fmt"

	"github.com/wtwingate/pokedex/internal/pokeapi"
)

func commandMapB(config *config) error {
	if config.Previous == nil {
		return errors.New("no previous locations to show")
	}
	locations := pokeapi.GetLocations(*config.Previous)

	config.Next = locations.Next
	config.Previous = locations.Previous

	results := locations.Results
	for _, v := range results {
		fmt.Println(v.Name)
	}
	return nil
}
