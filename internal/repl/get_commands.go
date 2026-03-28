package repl

import "github.com/dcanuto/pokedexcli/internal/pokecache"

type paginationConfig struct {
	next     *string
	previous *string
	cache    pokecache.Cache
}

type cliCommand struct {
	name        string
	description string
	Callback    func(*paginationConfig) error
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
	}
}
