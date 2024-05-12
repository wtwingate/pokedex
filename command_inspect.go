package main

import "fmt"

func commandInspect(cfg *Config, arg string) error {
	pokemon, ok := cfg.Pokedex[arg]
	if !ok {
		fmt.Printf("you haven't caught %v yet\n", arg)
		return nil
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, v := range pokemon.Stats {
		fmt.Printf(" - %v: %v\n", v.Stat.Name, v.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, v := range pokemon.Types {
		fmt.Printf(" - %v\n", v.Type.Name)
	}

	return nil
}
