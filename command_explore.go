package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, arguments []string) error {
	if len(arguments) != 1 {
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
