package pokeapi

import (
	"fmt"
	"io"
	"encoding/json"
	"net/http"
)

type Exploration struct {
	PokemonEncounters	[]PokemonEncounters	`json:"pokemon_encounters"`
}

type PokemonEncounters struct {
	Pokemon	Pokemon	`json:"pokemon"`
}

func CommandExplore(config *ConfigStruct, zone string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + zone

	pokemon, err := getEncounters(config, url)
	if err != nil {
		return err
	}

	for _, poke := range pokemon {
		fmt.Println(" -", poke)
	}
	return nil
}

func getEncounters(config *ConfigStruct, url string) ([]string, error) {
	cached, exists := config.Cache.Get(url)
	body := []byte{}
	if exists {
		body = cached
	} else {
		res, err := http.Get(url)
		if err != nil {
			return []string{}, err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
	}

	config.Cache.Add(url, body)

	var data Exploration
	err := json.Unmarshal(body, &data)
	if err != nil {
		return []string{}, err
	}
	
	var returnSlice []string
	for _, pokeEncounter := range data.PokemonEncounters {
		returnSlice = append(returnSlice, pokeEncounter.Pokemon.Name)
	}

	return returnSlice, nil
}