package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, arguments []string) error {
	if len(arguments) != 1 {
		return errors.New("only one pokemon name should be passed in with command")
	}
	pokemon, err := cfg.pokeApiClient.GetPokemonData(arguments[0])
	if err != nil {
		return err
	}
	// Attempt to help user catch specified pokemon
	fmt.Printf("Throwing a Pokeball at %s...\n", arguments[0])
	captureChance := rand.Intn(pokemon.BaseExperience)
	if captureChance >= (pokemon.BaseExperience / 2) {
		fmt.Printf("%s was caught!\n", arguments[0])
		fmt.Println("You may now inspect it with the inspect command")
		cfg.pokedex[arguments[0]] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", arguments[0])
	}
	return nil
}
