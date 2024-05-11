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
	defaultURL := baseURL + "location-area/"
	if pageURL == nil {
		pageURL = &defaultURL
	}

	locations := Locations{}

	if body, ok := c.cache.Get(*pageURL); ok {
		err := json.Unmarshal(body, &locations)
		if err != nil {
			return Locations{}, err
		}

		return locations, nil
	}

	body, err := c.getResource(*pageURL)
	if err != nil {
		return Locations{}, err
	}

	c.cache.Add(*pageURL, body)

	err = json.Unmarshal(body, &locations)
	if err != nil {
		return Locations{}, err
	}

	return locations, nil
}
