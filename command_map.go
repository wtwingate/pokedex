package main

import (
	"errors"
	"fmt"

	"github.com/wtwingate/pokedex/internal/pokeapi"
)

func commandMap(config *config) error {
	if config.Next == nil {
		return errors.New("no more locations to show")
	}
	locations, err := pokeapi.GetLocations(*config.Next)
	if err != nil {
		return err
	}

	config.Next = locations.Next
	config.Previous = locations.Previous

	results := locations.Results
	for _, v := range results {
		fmt.Println(v.Name)
	}
	return nil
}