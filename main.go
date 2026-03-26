package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dcanuto/pokedexcli/internal/repl"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := repl.PaginationConfig{}
	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		rawInput := scanner.Text()
		cleanedInput := repl.CleanInput(rawInput)

		commandName := cleanedInput[0]

		command, ok := repl.GetCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := command.Callback(&config)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
