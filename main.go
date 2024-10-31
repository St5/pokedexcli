package main

import (
	"time"

	"github.com/st5/pokedexcli/internal/pokeapi"
)

type config struct {
	poketapiClient pokeapi.Client
	nextLocUrl *string
	prevLocUrl *string
}

func main(){
	cfg := config{
		poketapiClient: pokeapi.NewClient(time.Minute),

	}
	
	startRepl(&cfg)
}