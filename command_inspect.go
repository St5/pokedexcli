package main

import "fmt"

func callbackInspect(cnf *config, params ...string) error {
	if len(params) < 1 {
		fmt.Println("Second param is requare. Please set pokemon name")
		return nil
	}

	for _, pockemonName := range params {

		PockemonStore, ok := cnf.caughtPokemons[pockemonName]
		if !ok {
			fmt.Printf("You haven`t caught that pokemon - %s \n", PockemonStore.Name)
			continue
		}

		fmt.Printf("-= %s =-\n", PockemonStore.Name)
		fmt.Printf("Height: %v \n", PockemonStore.Height)
		fmt.Printf("Weight: %v \n", PockemonStore.Weight)
		fmt.Println("Stats:")
		for _, stat := range PockemonStore.Stats {
			fmt.Printf(" - %s : %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, tt := range PockemonStore.Types {
			fmt.Printf(" - %s\n", tt.Type.Name)
		}

	}

	return nil
}