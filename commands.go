package main

import (
	"fmt"
	"os"

	"github.com/jdmukiibs/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(urlTracker *pokeApiUrlTracker) error
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
			description: "Displays the names of 20 location areas in the Pokemon world. Each subsequent call displays next 20 locations.",
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
	os.Exit(0)
	return nil
}

func commandMap(urlTracker *pokeApiUrlTracker) error {
	response, err := pokeapi.GetLocationAreas(urlTracker.next)
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}

func commandMapBack(urlTracker *pokeApiUrlTracker) error {
	response, err := pokeapi.GetLocationAreas(urlTracker.previous)
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}
