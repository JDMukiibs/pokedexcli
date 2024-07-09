package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

func main() {
	var exitPokedex bool
	log.SetPrefix("Pokedex: ")
	log.SetFlags(0)
	scanner := bufio.NewScanner(os.Stdin)

	for !exitPokedex {
		fmt.Print("pokedex > ")
		scanner.Scan()
		inputError := scanner.Err()
		for inputError != nil {
			log.Println("Error reading input command. Please enter a recognized command or enter 'help' to see available commands.")
			fmt.Print("pokedex > ")
			scanner.Scan()
			inputError = scanner.Err()
		}
		command := strings.TrimSpace(scanner.Text())
		if command == "exit" {
			exitPokedex = true
			continue
		}
		fmt.Println("You entered: ", command)
	}
}

func getCommandMap() map[string]cliCommand {
	return map[string]cliCommand {
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	commands := getCommandMap()
	for _, command := range commands {
		fmt.Println(command.name, ": ", command.description)
	}
	return nil
}

func commandExit() error {
	return errors.New("exit command has been triggered")
}
