package main

import "github.com/st5/pokedexcli/internal/pokeapi"

type config struct {
	poketapiClient pokeapi.Client
	nextLocUrl *string
	prevLocUrl *string
}

func main(){
	cfg := config{
		poketapiClient: pokeapi.NewClient(),

	}
	
	startRepl(&cfg)
}