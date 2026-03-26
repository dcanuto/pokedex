package repl

import (
	"bufio"
	"fmt"
	"os"
)

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	config := paginationConfig{}
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
			err := command.Callback(&config)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
