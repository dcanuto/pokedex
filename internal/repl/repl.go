package repl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type PaginationConfig struct {
	next     *string
	previous *string
}

type CliCommand struct {
	name        string
	description string
	Callback    func(*PaginationConfig) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
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

func commandHelp(config *PaginationConfig) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for commandName, command := range GetCommands() {
		fmt.Printf("%s: %s\n", commandName, command.description)
	}
	return nil
}

func commandExit(config *PaginationConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

type LocationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(config *PaginationConfig) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if config.next != nil {
		url = *config.next
	}
	return getLocationAreas(config, url)
}

func commandMapb(config *PaginationConfig) error {
	if config.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	url := *config.previous
	return getLocationAreas(config, url)
}

func getLocationAreas(config *PaginationConfig, url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	m := LocationArea{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	config.next = m.Next
	config.previous = m.Previous
	for _, result := range m.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func CleanInput(text string) []string {
	lowercased := strings.ToLower(text)
	return strings.Fields(lowercased)
}
