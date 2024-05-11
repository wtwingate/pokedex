package main

import (
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *Config, arg string) error {
	if len(arg) == 0 {
		fmt.Println("usage: catch <pokemon-name>")
		return nil
	}

	if _, ok := cfg.Pokedex[arg]; ok {
		fmt.Println("you've already captured", arg)
		return nil
	}

	pageURL := "https://pokeapi.co/api/v2/pokemon/" + arg

	data, err := cfg.Client.CatchPokemon(pageURL)
	if err != nil {
		return err
	}

	fmt.Printf("you throw a pokéball at %v", arg)
	for range 3 {
		time.Sleep(1 * time.Second)
		fmt.Printf(".")
	}
	fmt.Printf("\n")

	baseXP := data.BaseExperience
	if 100 > rand.Intn(baseXP) {
		cfg.Pokedex[arg] = data
		fmt.Printf("you caught %v!\n", arg)
	} else {
		fmt.Printf("%v escaped!\n", arg)
	}

	return nil
}
