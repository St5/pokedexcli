package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">>")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]

		if !ok {
			fmt.Println("invalid command")
			continue
		}

		params := cleaned[1:]

		err := command.callback(cfg, params...)

		if err != nil {
			fmt.Printf("Error %v", err)

		}
	}
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Help command",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit command",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "Explore the world of Pokemon. Displays the names of 20 location areas in the Pokemon world.",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Explore the world of Pokemon. Displays the prvious of the names of 20 location areas in the Pokemon world.",
			callback:    callbackMapB,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "It takes the name of a location area and show a list of all the Pokémon in a given area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokeman_name}",
			description: "Try to catch a pokemon",
			callback:    callbackCatch,
		},
		"inspect": {
			name: "inspect {poket_name}",
			description: "It print the name, height, weight, stats and type(s) of the Pokemon",
			callback: callbackInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "Print a list of all the names of the Pokemon the user",
			callback: callbackPokedex,
		},
	}
}
func cleanInput(str string) []string {
	lowered := strings.ToLower(str)

	words := strings.Fields(lowered)
	return words
}
