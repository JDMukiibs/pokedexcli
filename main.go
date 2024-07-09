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
	name        string
	description string
	callback    func() error
}

func main() {
	var exitPokedex error
	commands := getCommandMap()
	log.SetPrefix("Pokedex: ")
	log.SetFlags(0)
	scanner := bufio.NewScanner(os.Stdin)

	for exitPokedex == nil {
		input := getUserInput(scanner)
		userInput := strings.TrimSpace(input)
		command, ok := commands[userInput]
		for !ok {
			log.Println("Unrecognized command. Please enter a recognized command or enter 'help' to see available commands.")
			userInput = strings.TrimSpace(getUserInput(bufio.NewScanner(os.Stdin)))
			command, ok = commands[userInput]
		}
		switch command.name {
		case "help":
			command.callback()
			continue
		case "exit":
			exitPokedex = command.callback()
		}
	}
}

func getUserInput(scanner *bufio.Scanner) string {
	for {
		fmt.Print("pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			return input
		} else {
			inputError := scanner.Err()
			if inputError != nil {
				log.Println("Error reading input command. Please enter a recognized command or enter 'help' to see available commands.")
			}
		}
	}
}

func getCommandMap() map[string]cliCommand {
	return map[string]cliCommand{
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
	fmt.Print("\nUsage:\n\n")
	commands := getCommandMap()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	return errors.New("exit command has been triggered")
}
