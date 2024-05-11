package main

import "fmt"

func commandHelp(config *Config, arg string) error {
	for _, cmd := range getCommands() {
		fmt.Printf("%s\t%s\n", cmd.name, cmd.description)
	}
	return nil
}
