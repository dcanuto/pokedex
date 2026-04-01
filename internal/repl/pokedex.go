package repl

import (
	"fmt"
)

func commandPokedex(config *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for k := range config.pokedex {
		fmt.Printf("  - %s\n", k)
	}
	return nil
}
