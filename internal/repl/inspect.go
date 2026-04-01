package repl

import (
	"errors"
	"fmt"
)

func commandInspect(config *config, args ...string) error {
	if len(args) > 1 {
		return errors.New("inspect only accepts 1 argument")
	}
	pokemonName := args[0]
	pokemon, ok := config.pokedex[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
	} else {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  - %s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, pokemonType := range pokemon.Types {
			fmt.Printf("  - %s\n", pokemonType.Type.Name)
		}
	}
	return nil
}
