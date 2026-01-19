package pokeapi

import (
	"fmt"
	"io"
	"encoding/json"
	"net/http"
)

type ConfigStruct struct {
	NextURL		*string
	PrevURL		*string
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


func CommandMap(config *ConfigStruct) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if config.NextURL != nil {
		url = *config.NextURL
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	var data locationArea
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}
	
	locations := data.Results
	for _, area := range locations {
		fmt.Println(area.Name)
	}

	if data.Prev != nil {
		config.PrevURL = data.Prev
	}
	if data.Next != nil {
		config.NextURL = data.Next
	}
	
	return nil
}

func CommandMapB(config *ConfigStruct) error {
	if config.PrevURL == nil {
		fmt.Println("You're on the first page")
		return nil
	}
	url := config.PrevURL

	res, err := http.Get(*url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	var data locationArea
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}
	
	locations := data.Results
	for _, area := range locations {
		fmt.Println(area.Name)
	}

	if data.Prev != nil {
		config.PrevURL = data.Prev
	} else {
		config.PrevURL = nil
	}
	if data.Next != nil {
		config.NextURL = data.Next
	}
	
	return nil
}