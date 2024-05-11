package pokeapi

import (
	"encoding/json"
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

func (c *Client) GetLocations(pageURL *string) (Locations, error) {
	endpoint := baseURL + "location/"
	if pageURL != nil {
		endpoint = *pageURL
	}

	locations := Locations{}

	if body, ok := c.cache.Get(endpoint); ok {
		err := json.Unmarshal(body, &locations)
		if err != nil {
			return Locations{}, err
		}

		return locations, nil
	}

	body, err := c.getResource(endpoint)
	if err != nil {
		return Locations{}, err
	}

	c.cache.Add(endpoint, body)

	err = json.Unmarshal(body, &locations)
	if err != nil {
		return Locations{}, err
	}

	return locations, nil
}
