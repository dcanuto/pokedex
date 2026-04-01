package repl

import (
	"errors"
	"fmt"
)

func commandExplore(config *config, args ...string) error {
	if len(args) > 1 {
		return errors.New("explore only accepts 1 argument")
	}
	desiredLocation := args[0]
	location, err := config.client.GetLocation(desiredLocation)
	if err != nil {
		return err
	}

	for _, encounter := range location.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
