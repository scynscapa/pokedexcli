package pokeapi

import (
	"fmt"
	"io"
	"encoding/json"
	"net/http"
	"github.com/scynscapa/pokedexcli/internal/pokecache"
)

type ConfigStruct struct {
	NextURL		*string
	PrevURL		*string
	Cache		*pokecache.Cache
	Pokedex		map[string]Pokemon
}

type locationArea struct {
	Count		int					`json:"count"`
	Next		*string				`json:"next"`
	Prev		*string				`json:"previous"`
	Results		[]locationAreaList	`json:"results"`
}

type locationAreaList struct {
	Name		string	`json:"name"`
	Url			string	`json:"url"`
}

func CommandMap(config *ConfigStruct, arg string) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if config.NextURL != nil {
		url = *config.NextURL
	}

	locations, err := getLocations(config, url)
	if err != nil {
		return err
	}

	for _, area := range locations {
		fmt.Println(area)
	}

	return nil
}

func CommandMapB(config *ConfigStruct, arg string) error {
	if config.PrevURL == nil {
		fmt.Println("You're on the first page")
		return nil
	}
	url := config.PrevURL

	locations, err := getLocations(config, *url)
	if err != nil {
		return err
	}

	for _, area := range locations {
		fmt.Println(area)
	}
	
	return nil
}

func getLocations(config *ConfigStruct, url string) ([]string, error) {
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

	var data locationArea

	err := json.Unmarshal(body, &data)
	if err != nil {
		return []string{}, err
	}
	
	locations := data.Results

	config.Cache.Add(url, body)

	var returnSlice []string
	for _, area := range locations {
		returnSlice = append(returnSlice, area.Name)
	}

	if data.Prev != nil {
		config.PrevURL = data.Prev
	} else {
		config.PrevURL = nil
	}
	if data.Next != nil {
		config.NextURL = data.Next
	}

	return returnSlice, nil
}