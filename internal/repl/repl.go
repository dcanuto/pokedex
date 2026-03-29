package repl

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/dcanuto/pokedexcli/internal/pokecache"
)

func StartRepl() {
	const interval = 5 * time.Second

	scanner := bufio.NewScanner(os.Stdin)
	config := config{
		next:            nil,
		previous:        nil,
		desiredLocation: "",
		cache:           pokecache.NewCache(interval),
	}
	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		rawInput := scanner.Text()
		cleanedInput := cleanInput(rawInput)

		commandName := cleanedInput[0]

		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			if command.name == "explore" {
				config.desiredLocation = cleanedInput[1]
			}
			err := command.Callback(&config)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
