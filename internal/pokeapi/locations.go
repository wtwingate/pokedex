package pokeapi

import (
	"encoding/json"
	"log"
)

type Locations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocations(endpoint string) Locations {
	data := getResource(endpoint)
	locs := Locations{}
	err := json.Unmarshal(data, &locs)
	if err != nil {
		log.Fatal(err)
	}
	return locs
}
