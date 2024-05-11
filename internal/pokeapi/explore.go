package pokeapi

import (
	"encoding/json"
)

func (c *Client) ExploreLocation(pageURL string) (LocationAreaID, error) {
	explore := LocationAreaID{}

	if body, ok := c.cache.Get(pageURL); ok {
		err := json.Unmarshal(body, &explore)
		if err != nil {
			return LocationAreaID{}, err
		}

		return explore, nil
	}

	body, err := c.getResource(pageURL)
	if err != nil {
		return LocationAreaID{}, err
	}

	c.cache.Add(pageURL, body)

	err = json.Unmarshal(body, &explore)
	if err != nil {
		return LocationAreaID{}, err
	}

	return explore, nil
}
