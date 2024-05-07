package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	prompt := "Pokédex >>> "
	scanner := bufio.NewScanner(os.Stdin)

	for fmt.Print(prompt); scanner.Scan(); fmt.Print(prompt) {
		line := scanner.Text()
		if command, ok := cliCommandMap[line]; ok {
			command.callback()
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var cliCommandMap = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "Display a help message",
		callback:    commandHelp,
	},
	"exit": {
		name:        "exit",
		description: "Exit the program",
		callback:    commandExit,
	},
}

func commandHelp() error {
	fmt.Println()
	fmt.Println("No one can hear your cries for help...")
	fmt.Println()
	return nil
}

func commandExit() error {
	fmt.Println()
	fmt.Println("There is no escape!")
	fmt.Println()
	return nil
}
