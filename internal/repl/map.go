package repl

import (
	"fmt"
)

func commandMap(config *config, args ...string) error {
	locationAreas, err := config.client.GetLocationAreas(config.next)
	if err != nil {
		return err
	}

	config.next = locationAreas.Next
	config.previous = locationAreas.Previous
	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMapb(config *config, args ...string) error {
	if config.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	locationAreas, err := config.client.GetLocationAreas(config.previous)
	if err != nil {
		return err
	}

	config.next = locationAreas.Next
	config.previous = locationAreas.Previous
	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}
	return nil
}
