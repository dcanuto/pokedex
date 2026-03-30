package repl

import (
	"fmt"
)

func commandExplore(config *config, desiredLocation string) error {
	location, err := config.client.GetLocation(desiredLocation)
	if err != nil {
		return err
	}

	for _, encounter := range location.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
