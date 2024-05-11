package main

import (
	"fmt"
)

func commandMap(cfg *Config, arg string) error {
	data, err := cfg.Client.GetLocations(cfg.Next)
	if err != nil {
		return err
	}

	cfg.Next = data.Next
	cfg.Previous = data.Previous

	for _, v := range data.Results {
		fmt.Println(" -", v.Name)
	}
	return nil
}
