package main

import (
	"fmt"
	"time"
)

func callbackExplore(cnf *config, params ...string) error {

	if len(params) < 1 {
		fmt.Println("Second param is requare. Please set location name")
		return nil
	}
	for _, locationName := range params {
		start := time.Now()

		LocationInfo, err := cnf.poketapiClient.LocationInfo(locationName)

		if err != nil {
			return err
		}

		fmt.Println("Exploaring " + locationName)
		fmt.Println("Found pokemons:")

		for _, Pokemon := range LocationInfo.PokemonEncounters {
			fmt.Printf("  - %v \n", Pokemon.Pokemon.Name)
		}

		elapsed := time.Since(start)
		fmt.Printf("Request taken: %v \n", elapsed)

	}

	return nil
}
