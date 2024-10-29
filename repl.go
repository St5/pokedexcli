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
	callback func() error
}

func startRepl() {
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

		command.callback()
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
}
}
func cleanInput(str string) []string {
	lowered := strings.ToLower(str)

	words := strings.Fields(lowered)
	return words
}
