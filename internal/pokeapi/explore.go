package pokeapi

import (
	"encoding/json"
)

type Explore struct {
	// EncounterMethodRates []struct {
	// 	EncounterMethod struct {
	// 		Name string `json:"name"`
	// 		URL  string `json:"url"`
	// 	} `json:"encounter_method"`
	// 	VersionDetails []struct {
	// 		Rate    int `json:"rate"`
	// 		Version struct {
	// 			Name string `json:"name"`
	// 			URL  string `json:"url"`
	// 		} `json:"version"`
	// 	} `json:"version_details"`
	// } `json:"encounter_method_rates"`
	// GameIndex int `json:"game_index"`
	ID       int `json:"id"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name string `json:"name"`
	// Names []struct {
	// 	Language struct {
	// 		Name string `json:"name"`
	// 		URL  string `json:"url"`
	// 	} `json:"language"`
	// 	Name string `json:"name"`
	// } `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		// VersionDetails []struct {
		// 	EncounterDetails []struct {
		// 		Chance          int   `json:"chance"`
		// 		ConditionValues []any `json:"condition_values"`
		// 		MaxLevel        int   `json:"max_level"`
		// 		Method          struct {
		// 			Name string `json:"name"`
		// 			URL  string `json:"url"`
		// 		} `json:"method"`
		// 		MinLevel int `json:"min_level"`
		// 	} `json:"encounter_details"`
		// 	MaxChance int `json:"max_chance"`
		// 	Version   struct {
		// 		Name string `json:"name"`
		// 		URL  string `json:"url"`
		// 	} `json:"version"`
		// } `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) ExploreLocation(pageURL string) (Explore, error) {
	explore := Explore{}

	if body, ok := c.cache.Get(pageURL); ok {
		err := json.Unmarshal(body, &explore)
		if err != nil {
			return Explore{}, err
		}

		return explore, nil
	}

	body, err := c.getResource(pageURL)
	if err != nil {
		return Explore{}, err
	}

	c.cache.Add(pageURL, body)

	err = json.Unmarshal(body, &explore)
	if err != nil {
		return Explore{}, err
	}

	return explore, nil
}
