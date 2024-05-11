package main

import (
	"fmt"
)

func commandMap(cfg *Config) error {
	locations, err := cfg.Client.GetLocations(cfg.Next)
	if err != nil {
		return err
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

	for _, v := range locations.Results {
		fmt.Println("--", v.Name)
	}
	return nil
}
