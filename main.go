package main

import (
	"time"

	"github.com/st5/pokedexcli/internal/pokeapi"
)

type config struct {
	poketapiClient pokeapi.Client
	nextLocUrl     *string
	prevLocUrl     *string
	caughtPokemons map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		poketapiClient: pokeapi.NewClient(time.Minute),
		caughtPokemons: map[string]pokeapi.Pokemon{},
	}

	startRepl(&cfg)
}
