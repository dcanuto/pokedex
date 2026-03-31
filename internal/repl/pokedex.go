package repl

import (
	"fmt"
)

func commandPokedex(config *config, arg string) error {
	fmt.Println("Your Pokedex:")
	for k := range config.pokedex {
		fmt.Printf("  - %s\n", k)
	}
	return nil
}
