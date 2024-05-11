package main

import "os"

func commandExit(cfg *Config, arg string) error {
	os.Exit(0)
	return nil
}
