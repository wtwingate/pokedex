package main

import (
	"errors"
	"fmt"
)

func commandMapB(cfg *Config) error {
	if cfg.Previous == nil {
		return errors.New("no previous locations to show")
	}
	locations, err := cfg.Client.GetLocations(cfg.Previous)
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
