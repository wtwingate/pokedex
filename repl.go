package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func startREPL(config *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	fmt.Println("Welcome to the Pokédex CLI!")
	prompt := "Pokédex >>> "
	for fmt.Print(prompt); scanner.Scan(); fmt.Print(prompt) {
		line := scanner.Text()
		words := cleanInput(line)
		if cmd, ok := commands[words[0]]; ok {
			err := cmd.callback(config)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("unknown command")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func cleanInput(s string) []string {
	words := []string{}
	for _, w := range strings.Fields(s) {
		w = strings.ToLower(w)
		words = append(words, w)
	}
	return words
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
