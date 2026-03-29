package repl

import (
	"encoding/json"
	"fmt"

	"github.com/dcanuto/pokedexcli/internal/pokecache"
)

type locationResourceList struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(config *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if config.next != nil {
		url = *config.next
	}
	return getLocationAreas(config, url)
}

func commandMapb(config *config) error {
	if config.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	url := *config.previous
	return getLocationAreas(config, url)
}

func getLocationAreas(config *config, url string) error {
	body, err := pokecache.GetFromOrAddToCache(url, &config.cache)
	if err != nil {
		return err
	}

	m := locationResourceList{}
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
