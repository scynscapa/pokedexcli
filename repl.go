package main

import (
	"fmt"
	"strings"
	"os"
	"net/http"
	"encoding/json"
	"io"
)

type cliCommand struct {
	name		string
	description	string
	callback	func(config *configStruct) error
}

type configStruct struct {
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

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:		"exit",
			description:"Exit the Pokedex",
			callback:	commandExit,
		},
		"help": {
			name:		"help",
			description:"Displays a help message",
			callback:	commandHelp,
		},
		"map": {
			name:		"map",
			description:"Displays next page of locations",
			callback:	commandMap,
		},
		"mapb": {
			name:		"mapb",
			description:"Displays previous page of locations",
			callback:	commandMapB,
		},
	}
}


func cleanInput(text string) []string {
	lowered := strings.ToLower(text)

	trimmed := strings.TrimSpace(lowered)

	split := strings.Fields(trimmed)

	return split
}

func commandExit(config *configStruct) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *configStruct) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("usage:")
	fmt.Println("")
	
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

func commandMap(config *configStruct) error {
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

func commandMapB(config *configStruct) error {
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