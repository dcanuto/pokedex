package repl

import (
	"fmt"
	"os"
)

func commandExit(config *paginationConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
