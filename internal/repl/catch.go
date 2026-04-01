package repl

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(config *config, args ...string) error {
	if len(args) > 1 {
		return errors.New("catch only accepts 1 argument")
	}
	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := config.client.GetPokemon(name)
	if err != nil {
		return err
	}

	// Starter Pokemon (e.g., Squirtle) give around 60 experience,
	// while advanced ones (e.g., Mewtwo) give around 300. So this
	// gives about 95% chance to catch a starter,
	// and about 20% chance to catch something advanced.
	result := rand.Intn(pokemon.BaseExperience)
	if result > 60 {
		fmt.Printf("%s escaped!\n", name)
	} else {
		fmt.Printf("%s was caught!\n", name)
		fmt.Println("You may now inspect it with the inspect command.")
		_, ok := config.pokedex[name]
		if ok {
			fmt.Printf("%s already in Pokedex.\n", name)
		} else {
			fmt.Printf("Adding %s to Pokedex.\n", name)
			config.pokedex[name] = pokemon
		}
	}

	return nil
}
