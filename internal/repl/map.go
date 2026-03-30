package repl

import (
	"fmt"
)

func commandMap(config *config, location string) error {
	m, err := config.client.GetLocationAreas(config.next)
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

func commandMapb(config *config, location string) error {
	if config.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	m, err := config.client.GetLocationAreas(config.previous)
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
