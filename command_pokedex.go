package main

import "fmt"

func callbackPokedex(cnf *config, params ...string) error {
	if len(cnf.caughtPokemons) == 0 {
		fmt.Println("you haven`t gotten any pokemon yet!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pockemon := range cnf.caughtPokemons {
		fmt.Printf(" - %s\n", pockemon.Name)
	}

	return nil
}