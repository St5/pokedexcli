package main

import "fmt"

func callbackHelp(cnf *config, params ...string) error {
	fmt.Println("Welcome to the Pkedex")
	fmt.Println("List of available commands:")
	for _, cmd := range getCommands() {
		fmt.Printf(" - %s , %s \n", cmd.name, cmd.description)
	}
	fmt.Println(" ")
	return nil
}