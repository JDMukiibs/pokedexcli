package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jdmukiibs/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(urlTracker *pokeApiUrlTracker) error
}

// TODO (Joshua): Update struct to keep track of starting id for a location-area
type pokeApiUrlTracker struct {
	previous string
	next     string
}

func main() {
	var exitPokedex error
	commands := getCommandMap()
	urlTracker := initializePokeApiTracker()
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
			command.callback(&urlTracker)
			continue
		case "exit":
			exitPokedex = command.callback(&urlTracker)
			continue
		case "map":
			command.callback(&urlTracker)
			continue
		case "mapb":
			command.callback(&urlTracker)
			continue
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

func initializePokeApiTracker() pokeApiUrlTracker {
	return pokeApiUrlTracker{
		previous: "",
		next: "https://pokeapi.co/api/v2/location-area/",
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
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world.\nEach subsequent call displays next 20 locations.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations.",
			callback:    commandMapBack,
		},
	}
}

func commandHelp(urlTracker *pokeApiUrlTracker) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Print("\nUsage:\n\n")
	commands := getCommandMap()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit(urlTracker *pokeApiUrlTracker) error {
	return errors.New("exit command has been triggered")
}

func commandMap(urlTracker *pokeApiUrlTracker) error {
	response, err := pokeapi.GetLocations(urlTracker.next)
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}

func commandMapBack(urlTracker *pokeApiUrlTracker) error {
	response, err := pokeapi.GetLocations(urlTracker.previous)
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}
