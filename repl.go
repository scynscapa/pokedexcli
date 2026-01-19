package main

import (
	"fmt"
	"strings"
	"os"
	"github.com/scynscapa/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name		string
	description	string
	callback	func(config *pokeapi.ConfigStruct) error
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
			callback:	pokeapi.CommandMap,
		},
		"mapb": {
			name:		"mapb",
			description:"Displays previous page of locations",
			callback:	pokeapi.CommandMapB,
		},
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)

	trimmed := strings.TrimSpace(lowered)

	split := strings.Fields(trimmed)

	return split
}

func commandExit(config *pokeapi.ConfigStruct) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *pokeapi.ConfigStruct) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("usage:")
	fmt.Println("")
	
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
