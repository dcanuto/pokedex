package repl

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/dcanuto/pokedexcli/internal/pokeapi"
)

func StartRepl() {
	const interval = 5 * time.Second

	scanner := bufio.NewScanner(os.Stdin)
	config := config{
		next:     nil,
		previous: nil,
		client:   pokeapi.NewClient(interval),
		pokedex:  map[string]pokeapi.Pokemon{},
	}
	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		rawInput := scanner.Text()
		cleanedInput := cleanInput(rawInput)
		commandArgument := ""
		if len(cleanedInput) == 2 {
			commandArgument = cleanedInput[1]
		}

		commandName := cleanedInput[0]

		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := command.Callback(&config, commandArgument)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
