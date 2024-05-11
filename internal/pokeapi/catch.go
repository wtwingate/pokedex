package pokeapi

import "encoding/json"

func (c *Client) CatchPokemon(pageURL string) (PokemonID, error) {
	pokemon := PokemonID{}

	if body, ok := c.cache.Get(pageURL); ok {
		err := json.Unmarshal(body, &pokemon)
		if err != nil {
			return PokemonID{}, err
		}

		return pokemon, nil
	}

	body, err := c.getResource(pageURL)
	if err != nil {
		return PokemonID{}, err
	}

	c.cache.Add(pageURL, body)

	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return PokemonID{}, err
	}

	return pokemon, nil
}
