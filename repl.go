package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name string
	description string
	callback func(*config) error
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

		err := command.callback(cfg)

		if err != nil {
			fmt.Printf("Error %v", err)

		}
	}
}


func getCommands() map[string]cliCommand {

return map[string]cliCommand {
	"help": {
		name: "help",
		description: "Help command",
		callback: callbackHelp,
	},
	"exit": {
		name: "exit",
		description: "Exit command",
		callback: callbackExit,
	},
	"map": {
		name: "map",
		description: "Explore the world of Pokemon. Displays the names of 20 location areas in the Pokemon world.",
		callback: callbackMap,
	},
	"mapb": {
		name: "mapb",
		description: "Explore the world of Pokemon. Displays the prvious of the names of 20 location areas in the Pokemon world.",
		callback: callbackMapB,
	},
}
}
func cleanInput(str string) []string {
	lowered := strings.ToLower(str)

	words := strings.Fields(lowered)
	return words
}
