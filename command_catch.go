package main

import (
	"fmt"
	"math/rand"
	"time"
)

func callbackCatch(cnf *config, params ...string) error {
	if len(params) < 1 {
		fmt.Println("Second param is requare. Please set pokemon name")
		return nil
	}

	for _, pockemonName := range params {

		PockemonStore, ok := cnf.caughtPokemons[pockemonName]
		if ok {
			fmt.Printf("You already have that pokemon - %s \n", PockemonStore.Name)
			continue
		}
		start := time.Now()
		

		Pokemon, err := cnf.poketapiClient.GetPokemonInfo(pockemonName)

		if err != nil {
			fmt.Println("We don`t know such creature")
			return err
		}

		fmt.Printf("Throwing a Pokeball at %s (%d) \n", pockemonName, Pokemon.BaseExperience)

		elapsed := time.Since(start)
		fmt.Printf("Request taken: %v \n", elapsed)

		if !isCatch(Pokemon.BaseExperience) {
			fmt.Printf("%s escaped! \n", Pokemon.Name)
			continue
		}

		fmt.Printf("%s was caught! \n", Pokemon.Name)
		cnf.caughtPokemons[pockemonName] = Pokemon

	}

	return nil
}

func isCatch(n int) bool {
	return rand.Intn(n) <= 30
}
