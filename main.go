package main

import (
	"fmt"
	"bufio"
	"os"
	"github.com/scynscapa/pokedexcli/internal/pokeapi"
)

func main() {

	// create a scanner to read from stdin
	scanner := bufio.NewScanner(os.Stdin)

	conf := new(pokeapi.ConfigStruct)
	conf.NextURL = nil
	conf.PrevURL = nil

	for {
		// print prompt, scan for input
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		// clean input
		commandName := cleanInput(input)[0]

		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		command.callback(conf)
	}

}
