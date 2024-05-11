package main

import (
	"errors"
	"fmt"
)

func commandMapB(cfg *Config, arg string) error {
	if cfg.Previous == nil {
		return errors.New("no previous locations to show")
	}
	data, err := cfg.Client.GetLocations(cfg.Previous)
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
