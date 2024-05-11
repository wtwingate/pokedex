package pokeapi

import (
	"encoding/json"
)

func (c *Client) GetLocations(pageURL *string) (LocationArea, error) {
	defaultURL := baseURL + "location-area/"
	if pageURL == nil {
		pageURL = &defaultURL
	}

	locations := LocationArea{}

	if body, ok := c.cache.Get(*pageURL); ok {
		err := json.Unmarshal(body, &locations)
		if err != nil {
			return LocationArea{}, err
		}

		return locations, nil
	}

	body, err := c.getResource(*pageURL)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(*pageURL, body)

	err = json.Unmarshal(body, &locations)
	if err != nil {
		return LocationArea{}, err
	}

	return locations, nil
}
