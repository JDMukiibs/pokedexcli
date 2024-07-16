package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, arguments []string) error
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
		"explore": {
			name:        "explore",
			description: "Takes the name of a location area as an argument and lists all of the Pokémon in the requested area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Takes the name of a Pokemon as an argument and attempts to add them to your Pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Takes the name of a Pokemon as an argument and prints name, height, weight, stats, and type(s) of the Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Prints a list of all the names of the Pokemon you have caught. Takes no arguments",
			callback:    commandPokedex,
		},
	}
}

func commandHelp(config *config, arguments []string) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Print("\nUsage:\n\n")
	commands := getCommandMap()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit(config *config, arguments []string) error {
	os.Exit(0)
	return nil
}
