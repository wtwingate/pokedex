package main

import "fmt"

func commandPokedex(cfg *Config, arg string) error {
	if len(cfg.Pokedex) == 0 {
		fmt.Println("you haven't caught any Pokémon yet")
		return nil
	}

	fmt.Println("captured Pokémon:")
	for _, v := range cfg.Pokedex {
		fmt.Printf(" - %v\n", v.Name)
	}

	return nil
}
