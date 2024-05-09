package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL() {
	prompt := "Pokédex >>> "
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	endpoint := "https://pokeapi.co/api/v2/location/"
	config := config{&endpoint, nil}

	fmt.Println("Welcome to the Pokédex CLI!")
	for fmt.Print(prompt); scanner.Scan(); fmt.Print(prompt) {
		line := scanner.Text()
		words := cleanInput(line)
		if cmd, ok := commands[words[0]]; ok {
			err := cmd.callback(&config)
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
