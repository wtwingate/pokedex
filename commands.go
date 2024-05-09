package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/wtwingate/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     *string
	Previous *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "show the help menu",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "exit the program",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "show (next) 20 locations on map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "show (prev) 20 locations on map",
			callback:    commandMapB,
		},
	}
}

func commandHelp(config *config) error {
	for _, cmd := range getCommands() {
		fmt.Printf("%s\t%s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit(config *config) error {
	os.Exit(0)
	return nil
}

func commandMap(config *config) error {
	if config.Next == nil {
		return errors.New("no more locations to show")
	}
	locations := pokeapi.GetLocations(*config.Next)

	config.Next = locations.Next
	config.Previous = locations.Previous

	results := locations.Results
	for _, v := range results {
		fmt.Println(v.Name)
	}
	return nil
}

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
