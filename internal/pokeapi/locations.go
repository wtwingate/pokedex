package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return Locations{}, nil
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, err
	}

	if res.StatusCode > 299 {
		return Locations{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return Locations{}, err
	}

	locations := Locations{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return Locations{}, err
	}

	return locations, nil
}
