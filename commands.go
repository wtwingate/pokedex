package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokédex CLI!")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s:\t%s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func commandMap() error {
	return nil
}

func commandMapB() error {
	return nil
}
