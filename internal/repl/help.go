package repl

import "fmt"

func commandHelp(config *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for commandName, command := range getCommands() {
		fmt.Printf("%s: %s\n", commandName, command.description)
	}
	return nil
}
