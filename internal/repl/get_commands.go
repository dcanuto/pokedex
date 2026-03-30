package repl

import "github.com/dcanuto/pokedexcli/internal/pokeapi"

type config struct {
	next     *string
	previous *string
	client   pokeapi.Client
	pokedex  map[string]pokemon
}

type cliCommand struct {
	name        string
	description string
	Callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the next 20 locations",
			Callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 locations",
			Callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Display the Pokemon available in the provided area",
			Callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch the named Pokemon",
			Callback:    commandCatch,
		},
	}
}
