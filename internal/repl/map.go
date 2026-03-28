package repl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(config *paginationConfig) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if config.next != nil {
		url = *config.next
	}
	return getLocationAreas(config, url)
}

func commandMapb(config *paginationConfig) error {
	if config.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	url := *config.previous
	return getLocationAreas(config, url)
}

func getLocationAreas(config *paginationConfig, url string) error {
	body, exists := config.cache.Get(url)

	if !exists {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		config.cache.Add(url, body)
		fmt.Printf("Adding URL to cache: %s\n", url)
	} else {
		fmt.Printf("Retrieved areas from cache at URL: %s\n", url)
	}

	m := locationArea{}
	err := json.Unmarshal(body, &m)
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
