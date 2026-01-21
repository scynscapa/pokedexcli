package main

import (
	"fmt"
	"bufio"
	"os"
	"github.com/scynscapa/pokedexcli/internal/pokeapi"
	"github.com/scynscapa/pokedexcli/internal/pokecache"
)

func main() {

	// create a scanner to read from stdin
	scanner := bufio.NewScanner(os.Stdin)

	conf := new(pokeapi.ConfigStruct)
	conf.NextURL = nil
	conf.PrevURL = nil
	conf.Cache = pokecache.NewCache(5)
	conf.Pokedex = pokeapi.NewPokedex()

	for {
		// print prompt, scan for input
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		// clean input
		commandName := cleanInput(input)[0]

		argument := ""
		if len(cleanInput(input)) > 1 {
			argument = cleanInput(input)[1]
		}

		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(conf, argument)
		if err != nil {
			fmt.Println("ERROR: ", err)
		}
	}
}
