package main

import (
	"errors"
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

func commandMap(cfg *config, arguments []string) error {
	locationAreas, err := cfg.pokeApiClient.GetLocationAreas(cfg.nextUrl)
	if err != nil {
		return err
	}
	// Update urlTracker to have a new next and previous
	cfg.nextUrl = locationAreas.Next
	cfg.previousUrl = locationAreas.Previous
	// Print out our location areas
	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapBack(cfg *config, arguments []string) error {
	if cfg.previousUrl == nil {
		return errors.New("currently on first page of results. no previous results to show")
	}

	locationAreas, err := cfg.pokeApiClient.GetLocationAreas(cfg.previousUrl)
	if err != nil {
		return err
	}
	// Update urlTracker to have a new next and previous
	cfg.nextUrl = locationAreas.Next
	cfg.previousUrl = locationAreas.Previous
	// Print out our location areas
	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandExplore(cfg *config, arguments []string) error {
	if len(arguments) == 0 {
		return errors.New("name of a location area must be passed in with command")
	}
	locationAreaDetail, err := cfg.pokeApiClient.GetLocationAreaDetails(arguments[0])
	if err != nil {
		return err
	}
	// Print out pokemons found in the area
	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationAreaDetail.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
