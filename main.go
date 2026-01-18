package main

import (
	"fmt"
	"bufio"
	"os"
)





func main() {

	// create a scanner to read from stdin
	scanner := bufio.NewScanner(os.Stdin)

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

		command.callback()
	}

}
