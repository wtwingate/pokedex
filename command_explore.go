package main

import "fmt"

func commandExplore(cfg *Config, arg string) error {
	if len(arg) == 0 {
		fmt.Println("usage: explore <location-name>")
		return nil
	}

	pageURL := "https://pokeapi.co/api/v2/location-area/" + arg

	data, err := cfg.Client.ExploreLocation(pageURL)
	if err != nil {
		return err
	}

	fmt.Println("Pokémon found in this area:")
	for _, poke := range data.PokemonEncounters {
		fmt.Println(" -", poke.Pokemon.Name)
	}
	return nil
}
